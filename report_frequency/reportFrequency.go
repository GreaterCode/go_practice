/**
* Go语言词频统计，运行命令go run src/code/main.go test/words.txt
* @author unknown
* @since 2019-12-18
* 文件内容：
  hello tom glad to meet you
  yes me glad to meet you
  how are you
* 输出结果：
  Word        Frequency
  are   1
  glad  2
  hello 1
  how   1
  me    1
  meet  2
  to    2
  tom   1
  yes   1
  you   3
  Frequency → Words
  1 are, hello, how, me, tom, yes
  2 glad, meet, to
  3 you
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	// 命令行提示
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s file1 [ <file2> [... <fileN>]\n",filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	
	// 创建单词
	frequencyForWord := make(map[string] int) // 等价于map[string]int{}
	for _, filename := range commandLineFiles(os.Args[1:]) {
		updateFrequencies(filename,frequencyForWord)
	}
	// 打印单词=>频次
	reportByWords(frequencyForWord)

	// 翻转单词=>频次  为 频次=>单词多个
	wordsForFrequency := invertStringMap(frequencyForWord)

	//打印频次=>单词 多个
	reportByFrequency(wordsForFrequency)


}


func updateFrequencies(filename string,  frequencyForWard map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(filename);err != nil {
		log.Println("failed to open file:", err)
		return
	}
	defer file.Close()
	readerAndUpdateFrequencies(bufio.NewReader(file),frequencyForWard)
}

func readerAndUpdateFrequencies(reader *bufio.Reader, frequencyForWard map[string]int) {
	for  {
		line, err := reader.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)){
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				frequencyForWard[strings.ToLower(word)]+=1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file：", err)
			}
			break
		}
	}
}

func SplitOnNonLetters(s string) []string{
	notAllLeteers := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notAllLeteers)
}

/*
	因为Unix类系统的命令行工具会自动处理通配符（如*.text,会匹配1.txt、2.txt），而windows平台的命令行工具不支持通配符。
	为了保持平台之间一致性，使用commandLineFIles（）实现跨平台处理
*/
func commandLineFiles(files []string) []string {
	if runtime.GOARCH == "windoss" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err == nil {
				args = append(args, name) //invalid mode
			} else {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

func reportByWords(frequencyForWord map[string]int)  {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	/**
	* 经过排序之后，我们打印两列标题，第一个是word，为了能让Frequency最后一个字符y对齐
	* 需要在word 后打印一些空格，通过%*s可以实现的打印固定长度的空白，也可以使用%s来打印strings.Repeat(" ", gap) 返回的字符串
	 */
	gap := wordWidth + frequencyWidth - len("Word")- len("Frequency")
	fmt.Printf("Word %*s%s\n",gap," ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth,
			frequencyForWord[word])	}

}

/**
 * 首先创建一个空的映射，用来保存反转的结果，但是我们并不知道到期要保存多少项
 * 因此我们假设它和原来的映射容量一样大，然后简单地遍历原来的映射，将它的值作为键保存到翻转的映射里，并将键增加到对应的值里面去
 * 新的映射的值就是一个字符串的切片，即使原来的映射有多个键对应同一个值，也不会丢掉任何数据
 */
func invertStringMap(intForString map[string]int) map[int] []string  {
	stringForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringForInt[value] = append(stringForInt[value],key)
	}
	return stringForInt
}

/**
 * 首先创建一个切片用来保存频率，并按照频率升序排列，然后再计算需要容纳的最大长度并以此作为第一列的宽度，之后输出报告的标题
 * 最后，遍历输出所有的频率并按照字母升序输出对应的单词，如果一个频率有超过两个对应的单词，则单词之间以逗号分隔开
 */

func reportByFrequency(wordsForFrequency map[int][]string)  {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range frequencies {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))

	fmt.Println("Frequency-> words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words,","))
	}
}