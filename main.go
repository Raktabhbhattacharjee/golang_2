//  mic has a audio interface 

type Mic interface {
    CaptureAudio() string
}

//  differnt types of mic they all have their unique properties 
type SM7B struct{}

func (m SM7B) CaptureAudio() string {
    return "SM7B: Warm and deep — great for podcasts 🎙️"
}

// ----

type AT2020 struct{}

func (m AT2020) CaptureAudio() string {
    return "AT2020: Crisp and clear — great for vocals 🎙️"
}

// ----

type BlueYeti struct{}

func (m BlueYeti) CaptureAudio() string {
    return "Blue Yeti: Bright tone — great for streaming 🎙️"
}
func PlugInAndRecord(m Mic) {
    fmt.Println("🎛️  Focusrite is recording...")
    fmt.Println(m.CaptureAudio())
    fmt.Println("✅ Track saved!")
}
func main() {
    fmt.Println("--- Session 1 ---")
    PlugInAndRecord(SM7B{})

    fmt.Println("--- Session 2 ---")
    PlugInAndRecord(AT2020{})

    fmt.Println("--- Session 3 ---")
    PlugInAndRecord(BlueYeti{})
}
//  in interfaces we have the contracts or just the methood sgnatures so yea noting more than that 
// 