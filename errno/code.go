package errno

var (
	Success = &Errno{
		ErrorCode: "0",
		Note:      "Success",
	}
	Failure = &Errno{
		ErrorCode: "10000",
		Note:      "Failure",
	}
	FailureTimeFormat = &Errno{
		ErrorCode: "10001",
		Note:      "时间格式错误",
	}
	FailureParameter = &Errno{
		ErrorCode: "10002",
		Note:      "参数错误",
	}
	FailureJobExist = &Errno{
		ErrorCode: "10003",
		Note:      "任务已存在",
	}
	FailureNotFound = &Errno{
		ErrorCode: "10004",
		Note:      "查询不到数据",
	}
	FailureJsonUnmarshal = &Errno{
		ErrorCode: "10005",
		Note:      "jsonUnmarshal错误",
	}
	FailureJsonMarshal = &Errno{
		ErrorCode: "10006",
		Note:      "jsonMarshal错误",
	}
	FailureDbFind = &Errno{
		ErrorCode: "10007",
		Note:      "查询失败",
	}
	FailureDbUpdate = &Errno{
		ErrorCode: "10008",
		Note:      "更新失败",
	}

	FailureDbInsert = &Errno{
		ErrorCode: "10009",
		Note:      "插入失败",
	}
	FailureFileNotExist = &Errno{
		ErrorCode: "100010",
		Note:      "文件不存在",
	}
	FailureFileOpen = &Errno{
		ErrorCode: "100011",
		Note:      "文件打开句柄错误",
	}
	FailureFileCopy = &Errno{
		ErrorCode: "100012",
		Note:      "文件copy错误",
	}
	FailureFileExist = &Errno{
		ErrorCode: "100013",
		Note:      "文件已存在",
	}
	FailureFileMkdir = &Errno{
		ErrorCode: "100014",
		Note:      "创建目录失败",
	}
	FailureFileCreate = &Errno{
		ErrorCode: "100015",
		Note:      "创建文件失败",
	}
	FailureStrconv = &Errno{
		ErrorCode: "100016",
		Note:      "转换类型失败",
	}
	FailureUpdate = &Errno{
		ErrorCode: "100017",
		Note:      "更新失败",
	}
	FailureRpc = &Errno{
		ErrorCode: "100018",
		Note:      "RPC错误",
	}
	FailureFileWrite = &Errno{
		ErrorCode: "100019",
		Note:      "文件写入错误",
	}
	FailureFileRead = &Errno{
		ErrorCode: "100020",
		Note:      "文件读取错误",
	}
	FailureBase64Encode = &Errno{
		ErrorCode: "100021",
		Note:      "base64encode错误",
	}
	FailureBase64Decode = &Errno{
		ErrorCode: "100021",
		Note:      "base64decode错误",
	}
	FailureStatus = &Errno{
		ErrorCode: "100022",
		Note:      "status错误",
	}
	FailureFileCompress = &Errno{
		ErrorCode: "100023",
		Note:      "文件压缩错误",
	}
	FailureFileStat = &Errno{
		ErrorCode: "100024",
		Note:      "文件stat错误",
	}
	FailureExist = &Errno{
		ErrorCode: "100025",
		Note:      "数据已存在",
	}
	FailureNotExist = &Errno{
		ErrorCode: "100026",
		Note:      "数据不存在",
	}
	FailureDbDelete = &Errno{
		ErrorCode: "10027",
		Note:      "删除失败",
	}
	FailureProductNotFount = &Errno{
		ErrorCode: "10029",
		Note:      "未找到对应产品信息",
	}
)
