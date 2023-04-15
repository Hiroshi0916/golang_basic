package main

import (
	"fmt"
	"regexp"
)

// log
func main() {
	// log.SetOutput(os.Stdout)

	// log.Print("Log\n")
	// log.Println("Log2")
	// log.Printf("Log%d\n", 3)

	// log.Fatal("Log\n")
	// log.Fatal("Log2")
	// log.Fatal("Log%d\n", 3)

	// log.Panic("Log\n")
	// log.Panic("Log2")
	// log.Panic("Log%d\n", 3)

	//任意のファイルを作成し、出力先に指定
	//os.Create ファイルの作成
	// f, err := os.Create("test.log")
	// if err != nil {
	// 	return
	// }

	// //作成したio.Writer型のファイルを出力先に設定
	// log.SetOutput(f)
	// log.Println("ファイルに書き込む")

	// log.SetOutput(os.Stdout)
	// //ログのフォーマットを指定する
	// //標準のログフォーマット
	// log.SetFlags(log.LstdFlags)
	// log.Println("A")

	// //マイクロ秒を追加
	// log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	// log.Println("B")

	// //時刻とファイルの行番号(短縮系)
	// log.SetFlags(log.Ltime | log.Lshortfile)
	// log.Println("C")

	// log.SetFlags(log.LstdFlags)
	// // ログのプリフィックスを設定
	// log.SetPrefix("[LOG]")
	// log.Println("E")

	// //ロガーの生成
	// logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	// logger.Println("message")
	// log.Println("message")

	// //条件分岐、error終了
	// _, err := os.Open("fddafdsafa")
	// if err != nil {
	// 	//ログ出力
	// 	// log.Fatalln("Exit", err)
	// 	logger.Fatalln("Exit", err)
	// }

	//strconv

	// bt := true
	// fmt.Printf("%T\n", strconv.FormatBool(bt))

	//整数を文字列に。
	// i := strconv.FormatInt(-100, 10)
	// fmt.Printf("%v,%T\n", i, i)

	// //簡易的に変換できる
	// i2 := strconv.Itoa(100)
	// fmt.Printf("%v,%T\n", i2, i2)

	// //浮動小数型に変換する。
	// fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64))

	// fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64))

	// fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64))

	// fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))

	// fmt.Println(strconv.FormatFloat(123.456, 'g', -1, 64))

	// fmt.Println(strconv.FormatFloat(123456789.123, 'f', -1, 64))

	// fmt.Println(strconv.FormatFloat(123.456, 'g', 4, 64))

	// fmt.Println(strconv.FormatFloat(123456789.123, 'G', 8, 64))

	// bt1, _ := strconv.ParseBool("true")
	// fmt.Printf("%v,%T\n", bt1, bt1)
	// bt2, _ := strconv.ParseBool("1")
	// bt3, _ := strconv.ParseBool("t")
	// bt4, _ := strconv.ParseBool("T")
	// bt5, _ := strconv.ParseBool("TRUE")
	// bt6, _ := strconv.ParseBool("True")

	// fmt.Println(bt2, bt3, bt4, bt5, bt6)

	// //falseへ変換できる文字列
	// //2番目の戻り値はerror.
	// bf1, ok := strconv.ParseBool("false")
	// if ok != nil {
	// 	fmt.Println("Convert Error")
	// }
	// fmt.Printf("%v,%T\n", bf1, bf1)
	// bf2, _ := strconv.ParseBool("0")
	// bf3, _ := strconv.ParseBool("f")
	// bf4, _ := strconv.ParseBool("F")
	// bf5, _ := strconv.ParseBool("FALSE")
	// bf6, _ := strconv.ParseBool("False")

	// fmt.Println(bf2, bf3, bf4, bf5, bf6)

	// //文字列を整数型に変換
	// i3, _ := strconv.ParseInt("12345", 10, 0)
	// fmt.Printf("%v,%T\n", i3, i3)
	// i4, _ := strconv.ParseInt("-1", 10, 0)
	// fmt.Printf("%v,%T\n", i4, i4)

	// //簡易的に変換
	// i10, _ := strconv.Atoi("123")
	// fmt.Println(i10)

	// //文字列を浮動小数店へ。
	// fl1, _ := strconv.ParseFloat("3.14", 64)
	// fl2, _ := strconv.ParseFloat(".2", 64)
	// fl3, _ := strconv.ParseFloat("-2", 64)
	// fl4, _ := strconv.ParseFloat("1.2345e8", 64)
	// fl5, _ := strconv.ParseFloat("1.2345E8", 64)
	// fmt.Println(fl1, fl2, fl3, fl4, fl5)

	// //文字列を結合
	// s1 := strings.Join([]string{"A", "B", "C"}, ",")
	// s2 := strings.Join([]string{"A", "B", "C"}, "")
	// fmt.Println(s1, s2)

	// //文字列に含まれる部分文字列を検索
	// i1 := strings.Index("ABCDE", "A")
	// i2 := strings.Index("ABCDE", "D")
	// fmt.Println(i1, i2)

	// i3 := strings.LastIndex("ABCDABCD", "BC")
	// fmt.Println(i3)

	// i4 := strings.IndexAny("ABC", "ABC")

	// i5 := strings.LastIndexAny("ABC", "ABC")
	// fmt.Println(i4, i5)

	// b1 := strings.HasPrefix("ABC", "A")
	// b2 := strings.HasSuffix("ABC", "C")
	// fmt.Println(b1, b2)

	// b3 := strings.Contains("ABC", "B")
	// b4 := strings.ContainsAny("ABCDE", "BD")
	// fmt.Println(b3, b4)

	// i6 := strings.Count("ABCABC", "B")
	// i7 := strings.Count("ABCABC", "")
	// fmt.Println(i6, i7)

	// //文字列を繰り返して結合
	// s3 := strings.Repeat("ABC", 4)
	// s4 := strings.Repeat("ABC", 0)
	// fmt.Println(s3, s4)

	// //1は置換する回数の最大数
	// s5 := strings.Replace("AAAAA", "A", "B", 1)
	// //-1は該当する全てを置換
	// s6 := strings.Replace("AAAAA", "A", "B", -1)
	// fmt.Println(s5, s6)

	// //文字列を分割
	// s7 := strings.Split("A,B,C,D,E", ",")
	// fmt.Println(s7)

	// //,を残して分割
	// s8 := strings.SplitAfter("A,B,C,D,E", ",")
	// fmt.Println(s8)
	// //２分割
	// s9 := strings.SplitN("A,B,C,D,E", ",", 2)
	// fmt.Println(s9)
	// //,を残して４分割
	// s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
	// fmt.Println(s10)

	// //capital変換
	// s11 := strings.ToLower("ABC")
	// s12 := strings.ToLower("E")
	// s13 := strings.ToUpper("abc")
	// s14 := strings.ToUpper("e")
	// fmt.Println(s11, s12, s13, s14)

	// // 文字列の空白を取り除く
	// s15 := strings.TrimSpace("   -   ABC   -   ")
	// // 全角
	// s16 := strings.TrimSpace("      ABC      ")
	// fmt.Println(s15, s16)

	// //文字列からスペースで区切られたフィールドを取り出す
	// s18 := strings.Fields("a b c")
	// fmt.Println(s18)
	// fmt.Println(s18[1])

	// //bufio

	// //標準入力を行単位で。
	// scanner := bufio.NewScanner(os.Stdin)

	// //入力のスキャン
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "読み込みエラー", err)
	// }

	//ioutil
	// //入力全体を読み込む
	// f, _ := os.Open("foo.txt")
	// bs, _ := ioutil.ReadAll(f)

	// fmt.Println(string(bs))

	// if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
	// 	log.Fatalln(err)
	// }

	//regexp
	//正規表現

	// match, _ := regexp.MatchString("A", "ABC")
	// fmt.Println(match)

	// //Compile
	// re1, _ := regexp.Compile("A")
	// match = re1.MatchString("ABC")
	// fmt.Println(match)

	// //MustCompile
	// re2 := regexp.MustCompile("A")
	// match = re2.MatchString("ABC")
	// fmt.Println(match)

	// // 正規表現のフラグ
	// // i:大文字小文字を区別しない
	// // m マルチラインモード
	// // s .が\nにマッチ
	// // U 最小マッチへの変換

	// re3 := regexp.MustCompile(`(?i)abc`)
	// match := re3.MatchString("ABC")
	// fmt.Println(match)

	// //幅を持たない正規表現
	// re4 := regexp.MustCompile(`^ABC$`)
	// match := re4.MatchString("ABC")
	// fmt.Println(match)

	// match = re4.MatchString("   ABC   ")
	// fmt.Println(match)

	// //繰り返しを表す正規表現

	// re6 := regexp.MustCompile("a+b*")
	// fmt.Println(re6.MatchString("ab"))
	// fmt.Println(re6.MatchString("a"))

	// fmt.Println(re6.MatchString("aaaaaaabbbbbb"))

	// fmt.Println(re6.MatchString("b"))

	// //正規表現の文字クラス
	// re8 := regexp.MustCompile(`[XYZ]`)
	// fmt.Println(re8.MatchString("Y"))

	// re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	// fmt.Println(re9.MatchString("ABC"))
	// fmt.Println(re9.MatchString("abcdefg"))

	// re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	// fmt.Println(re10.MatchString("ABC"))
	// fmt.Println(re10.MatchString("あ"))

	// // 正規表現のグループ
	// re11 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	// fmt.Println(re11.MatchString("abcxyz"))

	// fmt.Println(re11.MatchString("ABCXYZ"))

	// fmt.Println(re11.MatchString("ABCxyz"))

	// fmt.Println(re11.MatchString("ABCabc"))

	// //FindString

	// re := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	// fs1 := re.FindString("AAAAABCXYZZZZ")
	// fmt.Println(fs1)
	// fs2 := re.FindAllString("ABCXYZABCXYZ", -1)
	// fmt.Println(fs2)

	// //正規表現のグループによるサブマッチ
	// re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	// s :=
	// 	`
	// 123-456-7889
	// 111-222-333
	// 556-787-899
	// `

	// ms := re1.FindAllStringSubmatch(s, -1)
	// fmt.Println(ms)

	// for _, v := range ms {
	// 	fmt.Println(v)
	// }

	// fmt.Println(ms[0][0])
	// fmt.Println(ms[0][1])
	// fmt.Println(ms[0][2])

	// //正規表現による文字列の置換
	// re1 := regexp.MustCompile(`\s+`)
	// fmt.Println(re1.ReplaceAllString("AAA BBB CCC, ", ""))

	// re2 := regexp.MustCompile(`、|。`)
	// fmt.Println(re2.ReplaceAllString("私はGolangを使用する、プログラマー", ""))

	//正規表現による文字列の分割

	re3 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	fmt.Println(re3.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1))

	re4 := regexp.MustCompile(`\s+`)
	fmt.Println(re4.Split("aaaaaaaa    bbbbbb   cccccc", -1))
}
