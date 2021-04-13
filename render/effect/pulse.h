#include <unistd.h>
#include <stdio.h>

#include <pulse/pulseaudio.h>

static pa_context *context = NULL;
static pa_stream *stream = NULL;
static char* device_name = NULL;
static char* device_description = NULL;
static pa_usec_t latency;
static unsigned channel_count;

extern void setLevels(int, float);

static void push_data(const float *d, unsigned samples) {
	while (samples >= channel_count) {

		for (unsigned c = 0; c < channel_count; c++) {
			setLevels(c, d[c]);
		}

		d += channel_count;
		samples -= channel_count;
	}
}

static void show_error(const char *txt) {
	fprintf(stderr, "%s: %s\n", txt, pa_strerror(pa_context_errno(context)));
}

static void stream_update_timing_info_callback(pa_stream *s, int success, void *d) {
	pa_usec_t t;
	int negative = 0;

	if (!success || pa_stream_get_latency(s, &t, &negative) < 0) {
		show_error("Failed to get latency information");
		return;
	}

	latency = negative ? 0 : t;
}

static void stream_read_callback(pa_stream *s, size_t l, void *d) {
	const void *p;

	if (pa_stream_peek(s, &p, &l) < 0) {
		fprintf(stderr, "pa_stream_peek() failed: %s", pa_strerror(pa_context_errno(context)));
		return;
	}

	push_data((const float*) p, l/sizeof(float));

	pa_stream_drop(s);
}

static void stream_state_callback(pa_stream *s, void *d) {
	pa_channel_map map;

	switch (pa_stream_get_state(s)) {
		case PA_STREAM_UNCONNECTED:
		case PA_STREAM_CREATING:
			break;

		case PA_STREAM_READY:
			map = *pa_stream_get_channel_map(s);
			channel_count = map.channels;

			pa_operation_unref(pa_stream_update_timing_info(stream, stream_update_timing_info_callback, NULL));
			break;

		case PA_STREAM_FAILED:
			show_error("Connection failed");
			break;

		case PA_STREAM_TERMINATED:
			break;
	}
}

static void create_stream(const char *name, const char *description, const pa_sample_spec ss, const pa_channel_map cmap) {
	char t[256];
	pa_sample_spec nss;

	device_name = pa_xstrdup(name);

	nss.format = PA_SAMPLE_FLOAT32;
	nss.rate = ss.rate;
	nss.channels = ss.channels;

	stream = pa_stream_new(context, "rgbd", &nss, &cmap);
	pa_stream_set_state_callback(stream, stream_state_callback, NULL);
	pa_stream_set_read_callback(stream, stream_read_callback, NULL);
	pa_stream_connect_record(stream, name, NULL, (enum pa_stream_flags) 0);
}

static void context_get_source_info_callback(pa_context *c, const pa_source_info *si, int is_last, void *d) {
	if (is_last < 0) {
		show_error("Failed to get source information");
		return;
	}

	if (!si)
		return;

	create_stream(si->name, si->description, si->sample_spec, si->channel_map);
}

static void context_get_sink_info_callback(pa_context *c, const pa_sink_info *si, int is_last, void *d) {
	if (is_last < 0) {
		show_error("Failed to get sink information");
		return;
	}

	if (!si)
		return;

	create_stream(si->monitor_source_name, si->description, si->sample_spec, si->channel_map);
}

static void context_get_server_info_callback(pa_context *c, const pa_server_info*si, void *d) {
	if (!si) {
		show_error("Failed to get server information");
		return;
	}


	if (!si->default_sink_name) {
		show_error("No default sink set.");
		return;
	}

	pa_operation_unref(pa_context_get_sink_info_by_name(c, si->default_sink_name, context_get_sink_info_callback, NULL));
}

static void context_state_callback(pa_context *c, void *d) {
	pa_context_state_t state = pa_context_get_state(c);
	switch (pa_context_get_state(c)) {
		case PA_CONTEXT_UNCONNECTED:
		case PA_CONTEXT_CONNECTING:
		case PA_CONTEXT_AUTHORIZING:
		case PA_CONTEXT_SETTING_NAME:
			break;

		case PA_CONTEXT_READY:
			pa_operation_unref(pa_context_get_sink_info_by_name(c, device_name, context_get_sink_info_callback, NULL));

			break;

		case PA_CONTEXT_FAILED:
			show_error("Connection failed");
			break;

		case PA_CONTEXT_TERMINATED:
			break;
	}
}

static void run_pulse() {
	pa_mainloop* mainloop = pa_mainloop_new();
	if (!mainloop) {
		fprintf(stderr, "pa_mainloop_new() failed.\n");
		return;
	}

	pa_mainloop_api* mainloop_api = pa_mainloop_get_api(mainloop);
	if (pa_signal_init(mainloop_api) != 0) {
		fprintf(stderr, "pa_signal_init() failed\n");
		return;
	}

	context = pa_context_new(mainloop_api, "PulseAudio Volume Meter");
	pa_context_set_state_callback(context, context_state_callback, NULL);
	if (pa_context_connect(context, NULL, PA_CONTEXT_NOAUTOSPAWN, NULL) != 0) {
		fprintf(stderr, "pa_context_connect() failed\n");
		return;
	}

	int ret;
	if (pa_mainloop_run(mainloop, &ret) < 0)
	{
		fprintf(stderr, "pa_mainloop_run() failed.\n");
		return;
	}
}
