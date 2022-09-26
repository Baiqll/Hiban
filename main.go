package main


import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"regexp"

)

func main() {

	var world_dict string

	flag.StringVar(&world_dict, "w", "", "url列表")

	// 解析命令行参数写入注册的flag里
	flag.Parse()


	// 如果管道有参数传递
	if has_stdin() {

		s := bufio.NewScanner(os.Stdin)

		for s.Scan() {
			
			get_all_version(s.Text())
			
		}

	}

	// 使用url字典列表
	if world_dict != "" {
		for _, url := range get_word_list(world_dict) {

			get_all_version(url)

		}
	}

}


// 获取所有版本信息
func get_all_version(url string){

	// 获取版本号
	// version_pattern := "/([0-9]{1,2}\.[0-9]{1,2}\.[0-9]{1,2})/"

	version_regexp := regexp.MustCompile("([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})")
	

	// 当url 只有出现一个版本号时
	if version := version_regexp.FindStringSubmatch(url); 3 > len(version) &&  len(version) > 0 {

		new_version_list := all_version(version[0])

		for _,new_version := range new_version_list{

			new_url := strings.Replace(url,version[0],new_version,-1)
			fmt.Println(new_url)
		}
		
	}

	//当url 没有版本号时
	fmt.Println(url)
	
}


// 获取 linux 管道传递的参数
func has_stdin() bool {

	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		return false
	}

	return true

}


// 返回字典列表
func get_word_list(file_path string) []string {

	f, err := os.Open(file_path)
	if err != nil {
		fmt.Println(err.Error())
	}
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(f)
	var dict_list []string
	for {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('\n')

		if errR == io.EOF {
			break
		}
		dict_list = append(dict_list, string(b))
	}

	return dict_list

}



func all_version(str string) []string {
	version := strings.Split(str, ".")
	res := []string{}
	x,_ := strconv.Atoi(version[0])
	y,_ :=  strconv.Atoi(version[1])
	z,_ :=  strconv.Atoi(version[2])

	for i :=0; i <=x+5; i++ {
		
		for j :=0; j <=y+5; j++ {
			for k :=0; k <=z+5; k++ {

				new_version := strings.Join([]string{strconv.Itoa(i),strconv.Itoa(j),strconv.Itoa(k)},".")
				res = append(res,new_version)
			}
		}
	}

	return res
}
