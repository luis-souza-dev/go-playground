package pkg

import (
	"fmt"
	"os/exec"
	"time"
)

func GetCode() {
	ch := make(chan string)
	_, err := exec.Command("which", "code").Output();
	var w1, w2 string
	fmt.Scanln(&w1, &w2);
	fmt.Println("works", w1)
	fmt.Println("works2", w2)

	if err != nil {
		fmt.Println("Vscode already installed");
	} else {
		go downloadCode(&ch);
	}
	time.Sleep(time.Second * 10);
	fmt.Println(<-ch)
}

func downloadCode(c *chan string) {
	fmt.Println("downloading .....");
	*c <- "hey there";
}