package main 

import (
	"context"
	"fmt"
	"regexp"

	"github.com/apache/beam/sdks/go/pkg/beam"
    "github.com/apache/beam/sdks/go/pkg/beam/runners/direct"
    "github.com/apache/beam/sdks/go/pkg/beam/transforms/stats"
)

func main() {
	beam.Init()
	p := beam.NewPipeline()
	s := p.Root()

	lines := beam.Create(s,
		"It was one hundred degrees",
		"As we sat beneath a willow tree",
		"Whose tears didn't care",
		"They just hung in the air",
		"And refused to fall, to fall-all-all-all",
		"And I knew I'd made a horrible call",
		"And now the state line felt like the Berlin Wall",
		"And there was no doubt about which side I was on",
		"Cause I built you a home in my heart",
		"With rotten wood it decayed from the start",
		"Cause you can't find nothing at all",
		"If there was nothing there all along",
		"No, you can't find nothing at all",
		"If there was nothing there all along",
		"I braced treacherous streets",
		"And kids strung out on homemade speed",
		"And we shared a bed in which I could not sleep",
		"At all",
		"Cause at night the sun in the tree",
		"Made the skyline look like crooked teeth",
		"In the mouth of a man who was devouring us both",
		"You're so cute when you're slurring your speech",
		"But they're closing the bar, and they want us to leave",
		"And you can't find nothing at all",
		"If there was nothing there all along",
		"No, you can't find nothing at all",
		"If there was nothing there all along",
		"I'm a war of head versus heart",
		"And it's always this way",
		"My head is weak",
		"My heart always speaks",
		"Before I know what it will say",
		"And you can't find nothing at all",
		"If there was nothing there all along",
		"No, you can't find nothing at all",
		"If there was nothing there all along",
		"And you can't find nothing at all",
		"If there was nothing there all along",
		"No, you can't find nothing at all",
		"If there was nothing there all along",
	)

	words := beam.ParDo(s, func(line string, emit func(string)) {
		for _, word := range wordRE.FindAllString(line, -1) {
			emit(word)
		}
	}, lines)

	counted := stats.Count(s, words)

	formatted := beam.ParDo(s, func(w string, c int) string {
		return fmt.Sprintf("%s: %v", w, c)
	}, counted)
	beam.ParDo(s, func(v string) { fmt.Println(v) }, formatted)

	direct.Execute(context.Background(), p)
}