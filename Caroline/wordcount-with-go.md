# WordCount using Beam Go
The following guide and demonstration will assume that you are using the Go Language playground or have installed Go and the necessary packages to run the Apache Beam SDK for Go and an IDE of your choice. You can find the playground here: [https://play.golang.org/](https://play.golang.org/)

To use beam within the playground, we must import it using the following code:

```Go
import (
    "context"
    "fmt"
    "regexp"

    "github.com/apache/beam/sdks/go/pkg/beam"
    "github.com/apache/beam/sdks/go/pkg/beam/runners/direct"
    "github.com/apache/beam/sdks/go/pkg/beam/transforms/stats"
)
```
Now we can get started with the Wordcount using Beam Go! We will be setting up a pipeline to read text and tokenize it into individual words. Once that is completed we will count the frequency inwhich the words appear.

## Create the Pipeline

A `Pipeline` object builds up a list of  the transformation that we will execute. We also need a scope, which is used to group the transformations into composite transforms.

```Go
func main() {
    beam.Init()
    //Create the pipeline object and root scope
    p := beam.NewPipeline()
    s := p.Root()
}
```

## Applying Pipeline Transformations

To perform a wordcount, our pipeline will contain several transforms so we can read the data into the pipeline -> transform it -> and get out our results. A little bit of lingo for ya, **composite transforms** are multiple nested transforms.

For each transformation, there is some kind of input data and output data. In the **Beam SDK** it is typically represented by the special class `PCollection`

```Go
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
```

We will use a **ParDo** transform to tokenize the text lines into individual words. The **ParDo** outputs a new `PCollection` where each element is an individual word in our text.

```Go
words := beam.ParDo(s, func(line string, emit func(string)) {
    for _, word := range wordRE.FindAllString(line, -1) {
        emit(word)
    }
}, lines)
```

The SDK-provided `Count` transform is a transform that takes a `PCollection` and returns a `PCollection of key/value pairs. Each key is the unique element and the value is how many times the key appeared in the input collection.

```Go
counted := stats.Count(s, words)
```

The next transform formats each of the key/value pairs of unique words and counts into a printable string. 

```Go
formatted := beam.ParDo(s, func(w string, c int) string {
    return fmt.Sprintf("%s: %v", w, c)
}, counted)
beam.ParDo(s, func(v string) {fmt.Println(v) }, formatted)
```

**Run the pipeline on the direct runner!**

```Go
direct.Execute(context.Background(), p)
```

## External Resources

- [A helpful walkthrough on the Go SDK](https://www.youtube.com/watch?v=95jis9rdrcg)
- [A browser playground fo Go](https://play.golang.org/p/6I1u_WM2oAr)
- [Beam's example on how to wordcount](https://beam.apache.org/get-started/wordcount-example/)
- [Information on what Beam SDK provides](https://beam.apache.org/get-started/beam-overview/)
