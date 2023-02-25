package main

/*
#include <stdio.h>
#include <setjmp.h>

jmp_buf env;
int ret = 0;

#define Try if ((ret=setjmp(env)) == 0)
#define Catch(e) else if (ret == (e))
#define Throw(e) longjmp(env, e)
#define Finally

char * eStrPtr2CStrPtr(char * str) {
	if (!str) {
		return NULL;
	}
	return str + 4;
}

int addX2(){
	int i = 2,j = 0;
	return i/j;
}

char * addX() {
    int idx = 0;
	char *c;
    Try {
		int k = 50;
		c = "化设计大撒是否";
        //printf("Try ...\n");
        //printf("Try222 ...\n");
        Throw(++idx);
		//return "";
    } Catch(1) {
        printf("Catch 1 ...\n");
		c = "哇户撒大苏打随风倒十分";
    }Finally {
		printf("滑稽666");
		//return "";
    }
	char *b;
	printf("滑稽888");
    return b;
}
*/
import "C"
import (
	"fmt"
	//sc "golang.org/x/example/encoding/simplifiedchinese"
)

//func CPtr2GoStr(str *C.char) string {
//	ptr := C.eStrPtr2CStrPtr(str)
//	if ptr != nil {
//		utf8str, _ := sc.GB18030.NewDecoder().String(C.GoString(ptr))
//		return utf8str
//	}
//	return ""
//}

func main() {
	huaji := Huaji()
	fmt.Println(huaji)
	fmt.Println("123")
}

func Huaji() (lala string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		lala = "滑稽了"
	}()
	str := C.GoString(C.addX())
	fmt.Println(str)
	fmt.Println("1234124123123")
	return lala
}
