package exiv2_test

import (
	"bytes"
	"testing"

	"github.com/robmurtha/expose/exiv2"
	"github.com/stretchr/testify/require"
)

func TestExpose(t *testing.T) {
	ex := exiv2.New(bytes.NewReader([]byte{}))
	fields, err := ex.Fields()
	require.NoError(t, err)
	require.Equal(t, 0, len(fields))

	for _, x := range groupTests {
		ex := exiv2.New(bytes.NewReader([]byte(x.data)))

		data, err := ex.Fields()
		require.NoError(t, err)
		require.Equal(t, x.count, len(data))
	}
}

var groupTests = []struct {
	data  string
	count int
}{
	{xmp, 64},
	{exif, 83},
	{iptc, 9},
}

var xmp = `video        FileSize                    XmpText     9  0.0515432
video        FileName                    XmpText    42  testdata/samsung_galaxy.mp4
video        MimeType                    XmpText    15  video/quicktime
video        MajorBrand                  XmpText    27  3GPP Media (.3GP) Release 4
video        MinorVersion                XmpText     3  768
video        CompatibleBrands            XmpSeq      3  3GPP Media (.3GP) Release 4, MP4 v1 [ISO 14496-1:ch13], 3GPP Media (.3GP) Release 6 Streaming Servers
video        MovieHeaderVersion          XmpText     1  0
video        DateUTC                     XmpText    10  3424280367
video        ModificationDate            XmpText    10  3424280367
video        TimeScale                   XmpText     4  1000
video        Duration                    XmpText     4  2268
video        PreferredRate               XmpText     1  1
video        PreferredVolume             XmpText     3  100
video        PreviewTime                 XmpText     1  0
video        PreviewDuration             XmpText     1  0
video        PosterTime                  XmpText     1  0
video        SelectionTime               XmpText     1  0
video        SelectionDuration           XmpText     1  0
video        CurrentTime                 XmpText     1  0
video        NextTrackID                 XmpText     1  3
video        TrackHeaderVersion          XmpText     1  0
video        TrackCreateDate             XmpText    10  3424280367
video        TrackModifyDate             XmpText    10  3424280367
video        TrackID                     XmpText     1  1
video        TrackDuration               XmpText     1  2
video        TrackLayer                  XmpText     1  0
video        TrackVolume                 XmpText     3  100
video        Width                       XmpText     3  320
video        Height                      XmpText     3  240
video        MediaHeaderVersion          XmpText     1  0
video        MediaCreateDate             XmpText    10  3424280367
video        MediaModifyDate             XmpText    10  3424280367
video        MediaTimeScale              XmpText     4  1000
video        MediaDuration               XmpText     1  2
video        MediaLangCode               XmpText     1  0
video        HandlerType                 XmpText    11  Video Track
video        GraphicsMode                XmpText     7  srcCopy
video        OpColor                     XmpText     1  0
video        URL                         XmpText     0
video        Codec                       XmpText     4  mp4v
video        SourceImageWidth            XmpText     3  320
video        SourceImageHeight           XmpText     3  240
video        XResolution                 XmpText     2  72
video        YResolution                 XmpText     2  72
video        Compressor                  XmpText     0
video        BitDepth                    XmpText     2  24
video        FrameRate                   XmpText     7  14.1093
audio        TrackHeaderVersion          XmpText     1  0
audio        TrackCreateDate             XmpText    10  3424280367
audio        TrackModifyDate             XmpText    10  3424280367
audio        TrackID                     XmpText     1  2
audio        TrackDuration               XmpText     1  2
audio        TrackLayer                  XmpText     1  0
audio        MediaHeaderVersion          XmpText     1  0
audio        MediaCreateDate             XmpText    10  3424280367
audio        MediaModifyDate             XmpText    10  3424280367
audio        MediaTimeScale              XmpText     4  1000
audio        MediaDuration               XmpText     1  2
audio        MediaLangCode               XmpText     1  0
audio        HandlerType                 XmpText    11  Audio Track
audio        Balance                     XmpText     1  0
audio        URL                         XmpText     0
audio        Compressor                  XmpText     4  mp4a
audio        ChannelType                 XmpText     1  1
audio        BitsPerSample               XmpText     2  16
audio        SampleRate                  XmpText     4  1000
video        AspectRatio                 XmpText     3  4:3`

var exif = `Image        Make                        Ascii       9  FUJIFILM
Image        Model                       Ascii       6  X100T
Image        Orientation                 Short       1  1
Image        XResolution                 Rational    1  72/1
Image        YResolution                 Rational    1  72/1
Image        ResolutionUnit              Short       1  2
Image        Software                    Ascii      29  Digital Camera X100T Ver1.00
Image        DateTime                    Ascii      20  2015:07:04 09:46:50
Image        YCbCrPositioning            Short       1  2
Image        Copyright                   Ascii       5
Image        ExifTag                     Long        1  352
Photo        ExposureTime                Rational    1  10/10000
Photo        FNumber                     Rational    1  450/100
Photo        ExposureProgram             Short       1  1
Photo        ISOSpeedRatings             Short       1  200
Photo        SensitivityType             Short       1  1
Photo        ExifVersion                 Undefined   4  48 50 51 48
Photo        DateTimeOriginal            Ascii      20  2015:07:04 09:46:50
Photo        DateTimeDigitized           Ascii      20  2015:07:04 09:46:50
Photo        ComponentsConfiguration     Undefined   4  1 2 3 0
Photo        CompressedBitsPerPixel      Rational    1  25/10
Photo        ShutterSpeedValue           SRational   1  1000/100
Photo        ApertureValue               Rational    1  430/100
Photo        BrightnessValue             SRational   1  966/100
Photo        ExposureBiasValue           SRational   1  0/100
Photo        MaxApertureValue            Rational    1  200/100
Photo        MeteringMode                Short       1  5
Photo        LightSource                 Short       1  0
Photo        Flash                       Short       1  16
Photo        FocalLength                 Rational    1  2300/100
Photo        MakerNote                   Undefined 772  (Binary value suppressed)
MakerNote    Offset                      Long        1  930
MakerNote    ByteOrder                   Ascii       3  II
Fujifilm     Version                     Undefined   4  48 49 51 48
Fujifilm     SerialNumber                Ascii      48  FFDT22404219     5933323031321411143C90201129C1
Fujifilm     Quality                     Ascii       8  NORMAL
Fujifilm     Sharpness                   Short       1  3
Fujifilm     WhiteBalance                Short       1  0
Fujifilm     Color                       Short       1  0
Fujifilm     FlashMode                   Short       1  2
Fujifilm     FlashStrength               SRational   1  0/100
Fujifilm     Macro                       Short       1  0
Fujifilm     FocusMode                   Short       1  0
Fujifilm     SlowSync                    Short       1  0
Fujifilm     PictureMode                 Short       1  768
Fujifilm     Continuous                  Short       1  1
Fujifilm     SequenceNumber              Short       1  4
Fujifilm     BlurWarning                 Short       1  0
Fujifilm     FocusWarning                Short       1  0
Fujifilm     ExposureWarning             Short       1  0
Fujifilm     DynamicRange                Short       1  1
Fujifilm     FilmMode                    Short       1  512
Fujifilm     DynamicRangeSetting         Short       1  0
Photo        FlashpixVersion             Undefined   4  48 49 48 48
Photo        ColorSpace                  Short       1  1
Photo        PixelXDimension             Long        1  1920
Photo        PixelYDimension             Long        1  1280
Photo        InteroperabilityTag         Long        1  1702
Iop          InteroperabilityIndex       Ascii       4  R98
Iop          InteroperabilityVersion     Undefined   4  48 49 48 48
Photo        FocalPlaneXResolution       Rational    1  820/1
Photo        FocalPlaneYResolution       Rational    1  820/1
Photo        FocalPlaneResolutionUnit    Short       1  3
Photo        SensingMethod               Short       1  2
Photo        FileSource                  Undefined   1  3
Photo        SceneType                   Undefined   1  1
Photo        CustomRendered              Short       1  0
Photo        ExposureMode                Short       1  1
Photo        WhiteBalance                Short       1  0
Photo        SceneCaptureType            Short       1  0
Photo        Sharpness                   Short       1  0
Photo        SubjectDistanceRange        Short       1  0
Image        PrintImageMatching          Undefined 106  80 114 105 110 116 73 77 0 48 50 53 48 0 0 3 0 2 0 1 0 0 0 3 0 34 0 0 0 1 1 0 0 0 0 9 17 0 0 16 39 0 0 11 15 0 0 16 39 0 0 151 5 0 0 16 39 0 0 176 8 0 0 16 39 0 0 1 28 0 0 16 39 0 0 94 2 0 0 16 39 0 0 139 0 0 0 16 39 0 0 203 3 0 0 16 39 0 0 229 27 0 0 16 39 0 0
Thumbnail    Compression                 Short       1  6
Thumbnail    Orientation                 Short       1  1
Thumbnail    XResolution                 Rational    1  72/1
Thumbnail    YResolution                 Rational    1  72/1
Thumbnail    ResolutionUnit              Short       1  2
Thumbnail    JPEGInterchangeFormat       Long        1  1856
Thumbnail    JPEGInterchangeFormatLength Long        1  8790
Thumbnail    YCbCrPositioning            Short       1  2
Image2       JPEGInterchangeFormat       Long        1  148
Image2       JPEGInterchangeFormatLength Long        1  905985`

var iptc = `Envelope     CharacterSet                String      3  G
Application2 RecordVersion               Short       1  4
Application2 DateCreated                 Date        8  2015-09-17
Application2 TimeCreated                 Time       11  11:35:39+00:00
Application2 DigitizationDate            Date        8  2015-09-17
Application2 DigitizationTime            Time       11  11:35:39+00:00
Application2 City                        String      5  Dover
Application2 ProvinceState               String      8  Delaware
Application2 CountryName                 String     13  United States`
