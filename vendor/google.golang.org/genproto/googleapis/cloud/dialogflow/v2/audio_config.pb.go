// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2/audio_config.proto

package dialogflow

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Audio encoding of the audio content sent in the conversational query request.
// Refer to the
// [Cloud Speech API
// documentation](https://cloud.google.com/speech-to-text/docs/basics) for more
// details.
type AudioEncoding int32

const (
	// Not specified.
	AudioEncoding_AUDIO_ENCODING_UNSPECIFIED AudioEncoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	AudioEncoding_AUDIO_ENCODING_LINEAR_16 AudioEncoding = 1
	// [`FLAC`](https://xiph.org/flac/documentation.html) (Free Lossless Audio
	// Codec) is the recommended encoding because it is lossless (therefore
	// recognition is not compromised) and requires only about half the
	// bandwidth of `LINEAR16`. `FLAC` stream encoding supports 16-bit and
	// 24-bit samples, however, not all fields in `STREAMINFO` are supported.
	AudioEncoding_AUDIO_ENCODING_FLAC AudioEncoding = 2
	// 8-bit samples that compand 14-bit audio samples using G.711 PCMU/mu-law.
	AudioEncoding_AUDIO_ENCODING_MULAW AudioEncoding = 3
	// Adaptive Multi-Rate Narrowband codec. `sample_rate_hertz` must be 8000.
	AudioEncoding_AUDIO_ENCODING_AMR AudioEncoding = 4
	// Adaptive Multi-Rate Wideband codec. `sample_rate_hertz` must be 16000.
	AudioEncoding_AUDIO_ENCODING_AMR_WB AudioEncoding = 5
	// Opus encoded audio frames in Ogg container
	// ([OggOpus](https://wiki.xiph.org/OggOpus)).
	// `sample_rate_hertz` must be 16000.
	AudioEncoding_AUDIO_ENCODING_OGG_OPUS AudioEncoding = 6
	// Although the use of lossy encodings is not recommended, if a very low
	// bitrate encoding is required, `OGG_OPUS` is highly preferred over
	// Speex encoding. The [Speex](https://speex.org/) encoding supported by
	// Dialogflow API has a header byte in each block, as in MIME type
	// `audio/x-speex-with-header-byte`.
	// It is a variant of the RTP Speex encoding defined in
	// [RFC 5574](https://tools.ietf.org/html/rfc5574).
	// The stream is a sequence of blocks, one block per RTP packet. Each block
	// starts with a byte containing the length of the block, in bytes, followed
	// by one or more frames of Speex data, padded to an integral number of
	// bytes (octets) as specified in RFC 5574. In other words, each RTP header
	// is replaced with a single byte containing the block length. Only Speex
	// wideband is supported. `sample_rate_hertz` must be 16000.
	AudioEncoding_AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE AudioEncoding = 7
)

var AudioEncoding_name = map[int32]string{
	0: "AUDIO_ENCODING_UNSPECIFIED",
	1: "AUDIO_ENCODING_LINEAR_16",
	2: "AUDIO_ENCODING_FLAC",
	3: "AUDIO_ENCODING_MULAW",
	4: "AUDIO_ENCODING_AMR",
	5: "AUDIO_ENCODING_AMR_WB",
	6: "AUDIO_ENCODING_OGG_OPUS",
	7: "AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE",
}

var AudioEncoding_value = map[string]int32{
	"AUDIO_ENCODING_UNSPECIFIED":            0,
	"AUDIO_ENCODING_LINEAR_16":              1,
	"AUDIO_ENCODING_FLAC":                   2,
	"AUDIO_ENCODING_MULAW":                  3,
	"AUDIO_ENCODING_AMR":                    4,
	"AUDIO_ENCODING_AMR_WB":                 5,
	"AUDIO_ENCODING_OGG_OPUS":               6,
	"AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE": 7,
}

func (x AudioEncoding) String() string {
	return proto.EnumName(AudioEncoding_name, int32(x))
}

func (AudioEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{0}
}

// Variant of the specified [Speech model][google.cloud.dialogflow.v2.InputAudioConfig.model] to use.
//
// See the [Cloud Speech
// documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models)
// for which models have different variants. For example, the "phone_call" model
// has both a standard and an enhanced variant. When you use an enhanced model,
// you will generally receive higher quality results than for a standard model.
type SpeechModelVariant int32

const (
	// No model variant specified. In this case Dialogflow defaults to
	// USE_BEST_AVAILABLE.
	SpeechModelVariant_SPEECH_MODEL_VARIANT_UNSPECIFIED SpeechModelVariant = 0
	// Use the best available variant of the [Speech
	// model][InputAudioConfig.model] that the caller is eligible for.
	//
	// Please see the [Dialogflow
	// docs](https://cloud.google.com/dialogflow-enterprise/docs/data-logging) for
	// how to make your project eligible for enhanced models.
	SpeechModelVariant_USE_BEST_AVAILABLE SpeechModelVariant = 1
	// Use standard model variant even if an enhanced model is available.  See the
	// [Cloud Speech
	// documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models)
	// for details about enhanced models.
	SpeechModelVariant_USE_STANDARD SpeechModelVariant = 2
	// Use an enhanced model variant:
	//
	// * If an enhanced variant does not exist for the given
	//   [model][google.cloud.dialogflow.v2.InputAudioConfig.model] and request language, Dialogflow falls
	//   back to the standard variant.
	//
	//   The [Cloud Speech
	//   documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models)
	//   describes which models have enhanced variants.
	//
	// * If the API caller isn't eligible for enhanced models, Dialogflow returns
	//   an error. Please see the [Dialogflow
	//   docs](https://cloud.google.com/dialogflow-enterprise/docs/data-logging)
	//   for how to make your project eligible.
	SpeechModelVariant_USE_ENHANCED SpeechModelVariant = 3
)

var SpeechModelVariant_name = map[int32]string{
	0: "SPEECH_MODEL_VARIANT_UNSPECIFIED",
	1: "USE_BEST_AVAILABLE",
	2: "USE_STANDARD",
	3: "USE_ENHANCED",
}

var SpeechModelVariant_value = map[string]int32{
	"SPEECH_MODEL_VARIANT_UNSPECIFIED": 0,
	"USE_BEST_AVAILABLE":               1,
	"USE_STANDARD":                     2,
	"USE_ENHANCED":                     3,
}

func (x SpeechModelVariant) String() string {
	return proto.EnumName(SpeechModelVariant_name, int32(x))
}

func (SpeechModelVariant) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{1}
}

// Gender of the voice as described in
// [SSML voice element](https://www.w3.org/TR/speech-synthesis11/#edef_voice).
type SsmlVoiceGender int32

const (
	// An unspecified gender, which means that the client doesn't care which
	// gender the selected voice will have.
	SsmlVoiceGender_SSML_VOICE_GENDER_UNSPECIFIED SsmlVoiceGender = 0
	// A male voice.
	SsmlVoiceGender_SSML_VOICE_GENDER_MALE SsmlVoiceGender = 1
	// A female voice.
	SsmlVoiceGender_SSML_VOICE_GENDER_FEMALE SsmlVoiceGender = 2
	// A gender-neutral voice.
	SsmlVoiceGender_SSML_VOICE_GENDER_NEUTRAL SsmlVoiceGender = 3
)

var SsmlVoiceGender_name = map[int32]string{
	0: "SSML_VOICE_GENDER_UNSPECIFIED",
	1: "SSML_VOICE_GENDER_MALE",
	2: "SSML_VOICE_GENDER_FEMALE",
	3: "SSML_VOICE_GENDER_NEUTRAL",
}

var SsmlVoiceGender_value = map[string]int32{
	"SSML_VOICE_GENDER_UNSPECIFIED": 0,
	"SSML_VOICE_GENDER_MALE":        1,
	"SSML_VOICE_GENDER_FEMALE":      2,
	"SSML_VOICE_GENDER_NEUTRAL":     3,
}

func (x SsmlVoiceGender) String() string {
	return proto.EnumName(SsmlVoiceGender_name, int32(x))
}

func (SsmlVoiceGender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{2}
}

// Audio encoding of the output audio format in Text-To-Speech.
type OutputAudioEncoding int32

const (
	// Not specified.
	OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_UNSPECIFIED OutputAudioEncoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	// Audio content returned as LINEAR16 also contains a WAV header.
	OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_LINEAR_16 OutputAudioEncoding = 1
	// MP3 audio.
	OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_MP3 OutputAudioEncoding = 2
	// Opus encoded audio wrapped in an ogg container. The result will be a
	// file which can be played natively on Android, and in browsers (at least
	// Chrome and Firefox). The quality of the encoding is considerably higher
	// than MP3 while using approximately the same bitrate.
	OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_OGG_OPUS OutputAudioEncoding = 3
)

var OutputAudioEncoding_name = map[int32]string{
	0: "OUTPUT_AUDIO_ENCODING_UNSPECIFIED",
	1: "OUTPUT_AUDIO_ENCODING_LINEAR_16",
	2: "OUTPUT_AUDIO_ENCODING_MP3",
	3: "OUTPUT_AUDIO_ENCODING_OGG_OPUS",
}

var OutputAudioEncoding_value = map[string]int32{
	"OUTPUT_AUDIO_ENCODING_UNSPECIFIED": 0,
	"OUTPUT_AUDIO_ENCODING_LINEAR_16":   1,
	"OUTPUT_AUDIO_ENCODING_MP3":         2,
	"OUTPUT_AUDIO_ENCODING_OGG_OPUS":    3,
}

func (x OutputAudioEncoding) String() string {
	return proto.EnumName(OutputAudioEncoding_name, int32(x))
}

func (OutputAudioEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{3}
}

// Instructs the speech recognizer how to process the audio content.
type InputAudioConfig struct {
	// Required. Audio encoding of the audio content to process.
	AudioEncoding AudioEncoding `protobuf:"varint,1,opt,name=audio_encoding,json=audioEncoding,proto3,enum=google.cloud.dialogflow.v2.AudioEncoding" json:"audio_encoding,omitempty"`
	// Required. Sample rate (in Hertz) of the audio content sent in the query.
	// Refer to
	// [Cloud Speech API
	// documentation](https://cloud.google.com/speech-to-text/docs/basics) for
	// more details.
	SampleRateHertz int32 `protobuf:"varint,2,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`
	// Required. The language of the supplied audio. Dialogflow does not do
	// translations. See [Language
	// Support](https://cloud.google.com/dialogflow-enterprise/docs/reference/language)
	// for a list of the currently supported language codes. Note that queries in
	// the same session do not necessarily need to specify the same language.
	LanguageCode string `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. The collection of phrase hints which are used to boost accuracy
	// of speech recognition.
	// Refer to
	// [Cloud Speech API
	// documentation](https://cloud.google.com/speech-to-text/docs/basics#phrase-hints)
	// for more details.
	PhraseHints []string `protobuf:"bytes,4,rep,name=phrase_hints,json=phraseHints,proto3" json:"phrase_hints,omitempty"`
	// Optional. Which variant of the [Speech model][google.cloud.dialogflow.v2.InputAudioConfig.model] to use.
	ModelVariant         SpeechModelVariant `protobuf:"varint,10,opt,name=model_variant,json=modelVariant,proto3,enum=google.cloud.dialogflow.v2.SpeechModelVariant" json:"model_variant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *InputAudioConfig) Reset()         { *m = InputAudioConfig{} }
func (m *InputAudioConfig) String() string { return proto.CompactTextString(m) }
func (*InputAudioConfig) ProtoMessage()    {}
func (*InputAudioConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{0}
}

func (m *InputAudioConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputAudioConfig.Unmarshal(m, b)
}
func (m *InputAudioConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputAudioConfig.Marshal(b, m, deterministic)
}
func (m *InputAudioConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputAudioConfig.Merge(m, src)
}
func (m *InputAudioConfig) XXX_Size() int {
	return xxx_messageInfo_InputAudioConfig.Size(m)
}
func (m *InputAudioConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_InputAudioConfig.DiscardUnknown(m)
}

var xxx_messageInfo_InputAudioConfig proto.InternalMessageInfo

func (m *InputAudioConfig) GetAudioEncoding() AudioEncoding {
	if m != nil {
		return m.AudioEncoding
	}
	return AudioEncoding_AUDIO_ENCODING_UNSPECIFIED
}

func (m *InputAudioConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

func (m *InputAudioConfig) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *InputAudioConfig) GetPhraseHints() []string {
	if m != nil {
		return m.PhraseHints
	}
	return nil
}

func (m *InputAudioConfig) GetModelVariant() SpeechModelVariant {
	if m != nil {
		return m.ModelVariant
	}
	return SpeechModelVariant_SPEECH_MODEL_VARIANT_UNSPECIFIED
}

// Description of which voice to use for speech synthesis.
type VoiceSelectionParams struct {
	// Optional. The name of the voice. If not set, the service will choose a
	// voice based on the other parameters such as language_code and gender.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The preferred gender of the voice. If not set, the service will
	// choose a voice based on the other parameters such as language_code and
	// name. Note that this is only a preference, not requirement. If a
	// voice of the appropriate gender is not available, the synthesizer should
	// substitute a voice with a different gender rather than failing the request.
	SsmlGender           SsmlVoiceGender `protobuf:"varint,2,opt,name=ssml_gender,json=ssmlGender,proto3,enum=google.cloud.dialogflow.v2.SsmlVoiceGender" json:"ssml_gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *VoiceSelectionParams) Reset()         { *m = VoiceSelectionParams{} }
func (m *VoiceSelectionParams) String() string { return proto.CompactTextString(m) }
func (*VoiceSelectionParams) ProtoMessage()    {}
func (*VoiceSelectionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{1}
}

func (m *VoiceSelectionParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoiceSelectionParams.Unmarshal(m, b)
}
func (m *VoiceSelectionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoiceSelectionParams.Marshal(b, m, deterministic)
}
func (m *VoiceSelectionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoiceSelectionParams.Merge(m, src)
}
func (m *VoiceSelectionParams) XXX_Size() int {
	return xxx_messageInfo_VoiceSelectionParams.Size(m)
}
func (m *VoiceSelectionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VoiceSelectionParams.DiscardUnknown(m)
}

var xxx_messageInfo_VoiceSelectionParams proto.InternalMessageInfo

func (m *VoiceSelectionParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VoiceSelectionParams) GetSsmlGender() SsmlVoiceGender {
	if m != nil {
		return m.SsmlGender
	}
	return SsmlVoiceGender_SSML_VOICE_GENDER_UNSPECIFIED
}

// Configuration of how speech should be synthesized.
type SynthesizeSpeechConfig struct {
	// Optional. Speaking rate/speed, in the range [0.25, 4.0]. 1.0 is the normal
	// native speed supported by the specific voice. 2.0 is twice as fast, and
	// 0.5 is half as fast. If unset(0.0), defaults to the native 1.0 speed. Any
	// other values < 0.25 or > 4.0 will return an error.
	SpeakingRate float64 `protobuf:"fixed64,1,opt,name=speaking_rate,json=speakingRate,proto3" json:"speaking_rate,omitempty"`
	// Optional. Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20
	// semitones from the original pitch. -20 means decrease 20 semitones from the
	// original pitch.
	Pitch float64 `protobuf:"fixed64,2,opt,name=pitch,proto3" json:"pitch,omitempty"`
	// Optional. Volume gain (in dB) of the normal native volume supported by the
	// specific voice, in the range [-96.0, 16.0]. If unset, or set to a value of
	// 0.0 (dB), will play at normal native signal amplitude. A value of -6.0 (dB)
	// will play at approximately half the amplitude of the normal native signal
	// amplitude. A value of +6.0 (dB) will play at approximately twice the
	// amplitude of the normal native signal amplitude. We strongly recommend not
	// to exceed +10 (dB) as there's usually no effective increase in loudness for
	// any value greater than that.
	VolumeGainDb float64 `protobuf:"fixed64,3,opt,name=volume_gain_db,json=volumeGainDb,proto3" json:"volume_gain_db,omitempty"`
	// Optional. An identifier which selects 'audio effects' profiles that are
	// applied on (post synthesized) text to speech. Effects are applied on top of
	// each other in the order they are given.
	EffectsProfileId []string `protobuf:"bytes,5,rep,name=effects_profile_id,json=effectsProfileId,proto3" json:"effects_profile_id,omitempty"`
	// Optional. The desired voice of the synthesized audio.
	Voice                *VoiceSelectionParams `protobuf:"bytes,4,opt,name=voice,proto3" json:"voice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SynthesizeSpeechConfig) Reset()         { *m = SynthesizeSpeechConfig{} }
func (m *SynthesizeSpeechConfig) String() string { return proto.CompactTextString(m) }
func (*SynthesizeSpeechConfig) ProtoMessage()    {}
func (*SynthesizeSpeechConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{2}
}

func (m *SynthesizeSpeechConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynthesizeSpeechConfig.Unmarshal(m, b)
}
func (m *SynthesizeSpeechConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynthesizeSpeechConfig.Marshal(b, m, deterministic)
}
func (m *SynthesizeSpeechConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynthesizeSpeechConfig.Merge(m, src)
}
func (m *SynthesizeSpeechConfig) XXX_Size() int {
	return xxx_messageInfo_SynthesizeSpeechConfig.Size(m)
}
func (m *SynthesizeSpeechConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SynthesizeSpeechConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SynthesizeSpeechConfig proto.InternalMessageInfo

func (m *SynthesizeSpeechConfig) GetSpeakingRate() float64 {
	if m != nil {
		return m.SpeakingRate
	}
	return 0
}

func (m *SynthesizeSpeechConfig) GetPitch() float64 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func (m *SynthesizeSpeechConfig) GetVolumeGainDb() float64 {
	if m != nil {
		return m.VolumeGainDb
	}
	return 0
}

func (m *SynthesizeSpeechConfig) GetEffectsProfileId() []string {
	if m != nil {
		return m.EffectsProfileId
	}
	return nil
}

func (m *SynthesizeSpeechConfig) GetVoice() *VoiceSelectionParams {
	if m != nil {
		return m.Voice
	}
	return nil
}

// Instructs the speech synthesizer on how to generate the output audio content.
type OutputAudioConfig struct {
	// Required. Audio encoding of the synthesized audio content.
	AudioEncoding OutputAudioEncoding `protobuf:"varint,1,opt,name=audio_encoding,json=audioEncoding,proto3,enum=google.cloud.dialogflow.v2.OutputAudioEncoding" json:"audio_encoding,omitempty"`
	// Optional. The synthesis sample rate (in hertz) for this audio. If not
	// provided, then the synthesizer will use the default sample rate based on
	// the audio encoding. If this is different from the voice's natural sample
	// rate, then the synthesizer will honor this request by converting to the
	// desired sample rate (which might result in worse audio quality).
	SampleRateHertz int32 `protobuf:"varint,2,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`
	// Optional. Configuration of how speech should be synthesized.
	SynthesizeSpeechConfig *SynthesizeSpeechConfig `protobuf:"bytes,3,opt,name=synthesize_speech_config,json=synthesizeSpeechConfig,proto3" json:"synthesize_speech_config,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *OutputAudioConfig) Reset()         { *m = OutputAudioConfig{} }
func (m *OutputAudioConfig) String() string { return proto.CompactTextString(m) }
func (*OutputAudioConfig) ProtoMessage()    {}
func (*OutputAudioConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{3}
}

func (m *OutputAudioConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputAudioConfig.Unmarshal(m, b)
}
func (m *OutputAudioConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputAudioConfig.Marshal(b, m, deterministic)
}
func (m *OutputAudioConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputAudioConfig.Merge(m, src)
}
func (m *OutputAudioConfig) XXX_Size() int {
	return xxx_messageInfo_OutputAudioConfig.Size(m)
}
func (m *OutputAudioConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputAudioConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutputAudioConfig proto.InternalMessageInfo

func (m *OutputAudioConfig) GetAudioEncoding() OutputAudioEncoding {
	if m != nil {
		return m.AudioEncoding
	}
	return OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_UNSPECIFIED
}

func (m *OutputAudioConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

func (m *OutputAudioConfig) GetSynthesizeSpeechConfig() *SynthesizeSpeechConfig {
	if m != nil {
		return m.SynthesizeSpeechConfig
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.dialogflow.v2.AudioEncoding", AudioEncoding_name, AudioEncoding_value)
	proto.RegisterEnum("google.cloud.dialogflow.v2.SpeechModelVariant", SpeechModelVariant_name, SpeechModelVariant_value)
	proto.RegisterEnum("google.cloud.dialogflow.v2.SsmlVoiceGender", SsmlVoiceGender_name, SsmlVoiceGender_value)
	proto.RegisterEnum("google.cloud.dialogflow.v2.OutputAudioEncoding", OutputAudioEncoding_name, OutputAudioEncoding_value)
	proto.RegisterType((*InputAudioConfig)(nil), "google.cloud.dialogflow.v2.InputAudioConfig")
	proto.RegisterType((*VoiceSelectionParams)(nil), "google.cloud.dialogflow.v2.VoiceSelectionParams")
	proto.RegisterType((*SynthesizeSpeechConfig)(nil), "google.cloud.dialogflow.v2.SynthesizeSpeechConfig")
	proto.RegisterType((*OutputAudioConfig)(nil), "google.cloud.dialogflow.v2.OutputAudioConfig")
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2/audio_config.proto", fileDescriptor_3ff9be2146363af6)
}

var fileDescriptor_3ff9be2146363af6 = []byte{
	// 898 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x6f, 0xe2, 0xc6,
	0x1b, 0xfe, 0xd9, 0x24, 0xfb, 0x53, 0xde, 0x90, 0xac, 0x77, 0x36, 0xcd, 0x3a, 0x69, 0x92, 0x92,
	0x64, 0x57, 0x62, 0x69, 0x0b, 0x2d, 0x2b, 0xf5, 0xd2, 0x93, 0xc1, 0x03, 0x58, 0x02, 0x63, 0xd9,
	0x40, 0xda, 0x5e, 0x46, 0x13, 0x7b, 0x30, 0x56, 0x8d, 0xc7, 0xc2, 0x86, 0xb6, 0x7b, 0xef, 0xb1,
	0x9f, 0xa1, 0x52, 0xd5, 0x53, 0xa5, 0x7e, 0xb6, 0x5e, 0x2b, 0xf5, 0x52, 0x79, 0x4c, 0x36, 0x2c,
	0xb0, 0x9c, 0x7a, 0xf3, 0x3c, 0xcf, 0xfb, 0xf7, 0x79, 0xdf, 0xf1, 0xc0, 0xe7, 0x3e, 0xe7, 0x7e,
	0xc8, 0x6a, 0x6e, 0xc8, 0xe7, 0x5e, 0xcd, 0x0b, 0x68, 0xc8, 0xfd, 0x71, 0xc8, 0x7f, 0xa8, 0x2d,
	0xea, 0x35, 0x3a, 0xf7, 0x02, 0x4e, 0x5c, 0x1e, 0x8d, 0x03, 0xbf, 0x1a, 0xcf, 0x78, 0xca, 0xd1,
	0x79, 0x6e, 0x5e, 0x15, 0xe6, 0xd5, 0x47, 0xf3, 0xea, 0xa2, 0x7e, 0x7e, 0xb1, 0x0c, 0x45, 0xe3,
	0xa0, 0x46, 0xa3, 0x88, 0xa7, 0x34, 0x0d, 0x78, 0x94, 0xe4, 0x9e, 0x37, 0x7f, 0xca, 0xa0, 0x18,
	0x51, 0x3c, 0x4f, 0xb5, 0x2c, 0x6a, 0x53, 0x04, 0x45, 0x16, 0x1c, 0xe7, 0x49, 0x58, 0xe4, 0x72,
	0x2f, 0x88, 0x7c, 0x55, 0x2a, 0x49, 0xe5, 0xe3, 0xfa, 0xeb, 0xea, 0x87, 0xf3, 0x54, 0x45, 0x00,
	0xbc, 0x74, 0xb0, 0x8f, 0xe8, 0xea, 0x11, 0x55, 0xe0, 0x59, 0x42, 0xa7, 0x71, 0xc8, 0xc8, 0x8c,
	0xa6, 0x8c, 0x4c, 0xd8, 0x2c, 0x7d, 0xab, 0xca, 0x25, 0xa9, 0xbc, 0x6f, 0x3f, 0xcd, 0x09, 0x9b,
	0xa6, 0xac, 0x93, 0xc1, 0xe8, 0x16, 0x8e, 0x42, 0x1a, 0xf9, 0x73, 0xea, 0x33, 0xe2, 0x72, 0x8f,
	0xa9, 0x85, 0x92, 0x54, 0x3e, 0xb0, 0x8b, 0x0f, 0x60, 0x93, 0x7b, 0x0c, 0x5d, 0x43, 0x31, 0x9e,
	0xcc, 0x68, 0xc2, 0xc8, 0x24, 0x88, 0xd2, 0x44, 0xdd, 0x2b, 0x15, 0xca, 0x07, 0xf6, 0x61, 0x8e,
	0x75, 0x32, 0x08, 0x39, 0x70, 0x34, 0xe5, 0x1e, 0x0b, 0xc9, 0x82, 0xce, 0x02, 0x1a, 0xa5, 0x2a,
	0x88, 0x26, 0xaa, 0xbb, 0x9a, 0x70, 0x62, 0xc6, 0xdc, 0x49, 0x2f, 0x73, 0x1b, 0xe5, 0x5e, 0x76,
	0x71, 0xba, 0x72, 0xba, 0xf9, 0x11, 0x4e, 0x46, 0x3c, 0x70, 0x99, 0xc3, 0x42, 0xe6, 0x66, 0x42,
	0x5a, 0x74, 0x46, 0xa7, 0x09, 0x42, 0xb0, 0x17, 0xd1, 0x29, 0x13, 0x42, 0x1d, 0xd8, 0xe2, 0x1b,
	0x75, 0xe1, 0x30, 0x49, 0xa6, 0x21, 0xf1, 0x59, 0xe4, 0xb1, 0x99, 0x68, 0xf7, 0xb8, 0xfe, 0xe9,
	0xce, 0xf4, 0xc9, 0x34, 0x14, 0xe1, 0xdb, 0xc2, 0xc5, 0x86, 0xcc, 0x3f, 0xff, 0xbe, 0xf9, 0x4b,
	0x82, 0x53, 0xe7, 0xa7, 0x28, 0x9d, 0xb0, 0x24, 0x78, 0xcb, 0xf2, 0x42, 0x97, 0xf3, 0xba, 0x85,
	0xa3, 0x24, 0x66, 0xf4, 0xfb, 0x20, 0xf2, 0x85, 0xbe, 0xa2, 0x0a, 0xc9, 0x2e, 0x3e, 0x80, 0x99,
	0xb6, 0xe8, 0x04, 0xf6, 0xe3, 0x20, 0x75, 0x27, 0xa2, 0x0e, 0xc9, 0xce, 0x0f, 0xe8, 0x25, 0x1c,
	0x2f, 0x78, 0x38, 0x9f, 0x32, 0xe2, 0xd3, 0x20, 0x22, 0xde, 0xbd, 0x50, 0x5b, 0xb2, 0x8b, 0x39,
	0xda, 0xa6, 0x41, 0xa4, 0xdf, 0xa3, 0xcf, 0x00, 0xb1, 0xf1, 0x98, 0xb9, 0x69, 0x42, 0xe2, 0x19,
	0x1f, 0x07, 0x21, 0x23, 0x81, 0xa7, 0xee, 0x0b, 0xcd, 0x95, 0x25, 0x63, 0xe5, 0x84, 0xe1, 0xa1,
	0x16, 0xec, 0x2f, 0xb2, 0x26, 0xd4, 0xbd, 0x92, 0x54, 0x3e, 0xac, 0x7f, 0xb1, 0xab, 0xe3, 0x6d,
	0x62, 0xda, 0xb9, 0xfb, 0xcd, 0xcf, 0x32, 0x3c, 0xeb, 0xcf, 0xd3, 0xb5, 0xe5, 0x1c, 0x7d, 0x60,
	0x39, 0x6b, 0xbb, 0xd2, 0xac, 0x84, 0xf9, 0x2f, 0x56, 0x34, 0x04, 0x35, 0x79, 0x37, 0x0a, 0x92,
	0x88, 0x59, 0x2c, 0x6f, 0xa4, 0xd0, 0xef, 0xb0, 0x5e, 0xdf, 0x39, 0xe6, 0xad, 0x63, 0xb4, 0x4f,
	0x93, 0xad, 0x78, 0xe5, 0x1f, 0x09, 0x8e, 0xde, 0x2b, 0x1d, 0x5d, 0xc1, 0xb9, 0x36, 0xd4, 0x8d,
	0x3e, 0xc1, 0x66, 0xb3, 0xaf, 0x1b, 0x66, 0x9b, 0x0c, 0x4d, 0xc7, 0xc2, 0x4d, 0xa3, 0x65, 0x60,
	0x5d, 0xf9, 0x1f, 0xba, 0x00, 0x75, 0x8d, 0xef, 0x1a, 0x26, 0xd6, 0x6c, 0xf2, 0xe5, 0x57, 0x8a,
	0x84, 0x5e, 0xc0, 0xf3, 0x35, 0xb6, 0xd5, 0xd5, 0x9a, 0x8a, 0x8c, 0x54, 0x38, 0x59, 0x23, 0x7a,
	0xc3, 0xae, 0x76, 0xa7, 0x14, 0xd0, 0x29, 0xa0, 0x35, 0x46, 0xeb, 0xd9, 0xca, 0x1e, 0x3a, 0x83,
	0x8f, 0x36, 0x71, 0x72, 0xd7, 0x50, 0xf6, 0xd1, 0xc7, 0xf0, 0x62, 0x8d, 0xea, 0xb7, 0xdb, 0xa4,
	0x6f, 0x0d, 0x1d, 0xe5, 0x09, 0x7a, 0x0d, 0xaf, 0xd6, 0x48, 0xc7, 0xc2, 0xf8, 0x1b, 0x72, 0x67,
	0x0c, 0x3a, 0xa4, 0x83, 0x35, 0x1d, 0xdb, 0xa4, 0xf1, 0xed, 0x00, 0x2b, 0xff, 0xaf, 0x2c, 0x00,
	0x6d, 0xde, 0x4a, 0xf4, 0x12, 0x4a, 0x99, 0x47, 0xb3, 0x43, 0x7a, 0x7d, 0x1d, 0x77, 0xc9, 0x48,
	0xb3, 0x0d, 0xcd, 0x1c, 0xac, 0xe9, 0x70, 0x0a, 0x68, 0xe8, 0x60, 0xd2, 0xc0, 0xce, 0x80, 0x68,
	0x23, 0xcd, 0xe8, 0x6a, 0x8d, 0x2e, 0x56, 0x24, 0xa4, 0x40, 0x31, 0xc3, 0x9d, 0x81, 0x66, 0xea,
	0x9a, 0xad, 0x2b, 0xf2, 0x03, 0x82, 0xcd, 0x8e, 0x66, 0x36, 0xb1, 0xae, 0x14, 0x2a, 0xbf, 0x48,
	0xf0, 0x74, 0xed, 0x3e, 0xa2, 0x6b, 0xb8, 0x74, 0x9c, 0x5e, 0x97, 0x8c, 0xfa, 0x46, 0x13, 0x93,
	0x36, 0x36, 0xb3, 0x3a, 0xdf, 0x4f, 0x79, 0x0e, 0xa7, 0x9b, 0x26, 0x3d, 0x4d, 0xa4, 0xbd, 0x00,
	0x75, 0x93, 0x6b, 0x61, 0xc1, 0xca, 0xe8, 0x12, 0xce, 0x36, 0x59, 0x13, 0x0f, 0x07, 0xb6, 0xd6,
	0x55, 0x0a, 0x95, 0xdf, 0x25, 0x78, 0xbe, 0x65, 0x8d, 0xd1, 0x2b, 0xb8, 0xee, 0x0f, 0x07, 0xd6,
	0x70, 0x40, 0x76, 0xae, 0xc4, 0x2d, 0x7c, 0xb2, 0xdd, 0x6c, 0x75, 0x33, 0x2e, 0xe1, 0x6c, 0xbb,
	0x51, 0xcf, 0x7a, 0xa3, 0xc8, 0xe8, 0x06, 0xae, 0xb6, 0xd3, 0xef, 0x26, 0x5b, 0x68, 0xfc, 0x2a,
	0xc1, 0x95, 0xcb, 0xa7, 0x3b, 0xd6, 0xbf, 0xa1, 0xac, 0x5c, 0x67, 0x2b, 0x7b, 0x85, 0x2c, 0xe9,
	0x3b, 0x7d, 0x69, 0xef, 0xf3, 0xec, 0x37, 0x5f, 0xe5, 0x33, 0xbf, 0xe6, 0xb3, 0x48, 0xbc, 0x51,
	0xb5, 0x9c, 0xa2, 0x71, 0x90, 0x6c, 0x7b, 0x0f, 0xbf, 0x7e, 0x3c, 0xfd, 0x2d, 0x49, 0xbf, 0xc9,
	0xb2, 0xde, 0xfa, 0x43, 0x3e, 0x6f, 0xe7, 0xe1, 0x9a, 0x22, 0xbd, 0xfe, 0x98, 0x7e, 0x54, 0xbf,
	0x7f, 0x22, 0xa2, 0xbe, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x15, 0xde, 0x52, 0x01, 0x64, 0x07,
	0x00, 0x00,
}
