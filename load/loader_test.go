package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"testing"
)

var FakeMasterPlayList = parse.MasterPlaylist{
	Path: "testMasterPlaylist.m3u8",
	Tags: []string{
		"#EXTM3U",
		"#EXT-X-VERSION:4",
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720",
	},
	MediaPlaylists: []parse.MediaPlaylist{
		{
			Path: "testMediaPlaylist.m3u8",
			Tags: []string{
				"#EXTM3U",
				"#EXT-X-VERSION:4",
				"#EXT-X-PLAYLIST-TYPE:VOD",
				"#EXT-X-INDEPENDENT-SEGMENTS",
				"#EXT-X-TARGETDURATION:5",
				"#EXT-X-MEDIA-SEQUENCE:0",
				"#EXT-X-DISCONTINUITY-SEQUENCE:0",
				"#EXTINF:3,",
				"#EXTINF:4,",
				"#EXTINF:5,",
				"#EXTINF:3,",
				"#EXTINF:4,",
				"#EXTINF:5,",
				"#EXTINF:3,",
				"#EXTINF:4,",
				"#EXTINF:5,",
				"#EXT-X-ENDLIST",
			},
			Segments: []parse.Segment{
				{Path: "segment-0.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-1.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-2.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-3.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-4.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-5.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-6.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-7.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-8.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
			},
			TotalDurationMs:         36 * 1000,
			TotalDiscontinuityCount: 0,
		},
	},
}

func TestVodLoadMasterPlaylist(t *testing.T) {
	loader := NewDefaultLoader(FakeMasterPlayList)
	actual, err := loader.LoadMasterPlaylist()
	if err != nil {
		t.Fatalf("error happened: %#v", err)
	}
	masterPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720\n" +
		"0/playlist.m3u8\n"
	if masterPlaylist != string(actual) {
		t.Errorf("exspected: %v, actual: %v", masterPlaylist, string(actual))
	}
}
