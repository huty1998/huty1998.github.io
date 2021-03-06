package main

import "fmt"

func main() {
	// bytes := []byte(`nvv4l2camerasrc name=nvv4l2camerasrc0 device=/dev/video0 ! capsfilter name=capsfilter0 caps="video/x-raw(memory:NVMM),format=(string)UYVY,width=1920,height=1080,framerate=30/1" ! tee name=t0 t0.src_0 ! queue name=queuet00 ! nvvidconv name=nvvidconvt00 ! capsfilter name=capsfiltert00 caps="video/x-raw(memory:NVMM),format=RGBA" t0.src_1 ! queue name=queuet01 ! mediamixer.sink_ ! eswcloseup0.sink_0 eswcloseup name=eswcloseup0 ! tee name=DtTeacherAi DtTeacherAi.src_0 ! queue name=queue0 ! mediamixer.sink_ eswmediamixer name=mediamixer config-file-path=/home/user/eswin/director/directorSpec.json ! videorate name=videorate0 ! capsfilter name=capsfilter1 caps="video/x-raw(memory:NVMM)" ! tee name=Dt Dt.src_1 ! queue name=queuepv0 ! nvvidconv name=nvvidconvpv0 ! capsfilter name=capsfilter2 caps="video/x-raw" ! xvimagesink name=xvimagesinkpv0 sync=false t0.src_2 ! queue name=queuet02 ! nvvidconv name=nvvidconvpvt00 ! capsfilter name=capsfilterpvt00 caps="video/x-raw,format=(string)UYVY,width=1920,height=1080,framerate=30/1" ! xvimagesink name=xvimagesinkt00 sync=false t0.src_3 ! queue name=queuet03 ! videorate name=videoratet00 max-rate=10 ! nvvidconv name=nvvidconvt01 ! capsfilter name=capsfiltert01 caps="video/x-raw,format=(string)UYVY,width=1920,height=1080" ! nvvideoconvert name=nvvideoconvertt00 ! capsfilter name=capsfiltert02 caps="video/x-raw(memory:NVMM),width=1920,height=1080,format=RGBA" ! m0.sink_0 nvstreammux name=m0 batch-size=1 batched-push-timeout=40000 height=1080 width=1920 ! nvinfer name=nvinfert0 config-file-path=/usr/lib/eswin/aarch64/ai-config/config_infer_primary_yoloV5.txt ! nvtracker name=nvtrackert0 enable-batch-process=1 gpu-id=0 ll-lib-file=/usr/lib/eswin/aarch64/ai-config/libnvds_nvmultiobjecttracker.so tracker-height=384 tracker-width=640 ! eswfilterobject name=eswfilterobjectt0 config-path=/home/user/.config/ai_config/platformestimate.txt ! nvinfer name=nvinfert1 config-file-path=/usr/lib/eswin/aarch64/ai-config/config_infer_secondary_pose17.txt ! eswposepostprocess name=eswposekeypointt0 ! eswteacheractrec name=eswteacheractrect0 ! nvvideoconvert name=nvvideoconvertt0 ! eswteacherdatacreator name=eswteacherdatacreatort0 ! eswteachersmartdirector name=eswteachersmartdirectort0 config-path=/usr/lib/eswin/aarch64/ai-config/config_teachersmartdirector.txt ! eswcloseup0.sink_1 v4l2src name=v4l2src0 device=/dev/video3 ! capsfilter name=capsfilter3 caps="video/x-raw,format=(string)YUY2,width=1920,height=1080,framerate=30/1" ! tee name=t1 t1.src_0 ! queue name=queuet10 ! nvvidconv name=nvvidconvt10 ! capsfilter name=capsfiltert10 caps="video/x-raw(memory:NVMM),format=RGBA" t1.src_1 ! queue name=queuet11 ! mediamixer.sink_ ! eswcloseup1.sink_0 eswcloseup name=eswcloseup1 ! tee name=DtStudentAi DtStudentAi.src_0 ! queue name=queue1 ! mediamixer.sink_ t1.src_2 ! queue name=queuet12 ! xvimagesink name=xvimagesinkt10 sync=false t1.src_3 ! queue name=queuet13 ! nvvideoconvert name=nvvideoconvertt10 ! capsfilter name=capsfiltert11 caps="video/x-raw(memory:NVMM),width=1920,height=1080,format=RGBA,framerate=30/1" ! m1.sink_0 nvstreammux name=m1 batch-size=1 batched-push-timeout=40000 height=1080 width=1920 ! nvinfer name=nvinfers0 config-file-path=/usr/lib/eswin/aarch64/ai-config/config_infer_primary_yoloV5n.txt ! nvtracker name=nvtrackers0 enable-batch-process=1 gpu-id=0 ll-config-file=/usr/lib/eswin/aarch64/ai-config/config_tracker_student.yml ll-lib-file=/usr/lib/eswin/aarch64/ai-config/libnvds_nvmultiobjecttracker.so tracker-height=384 tracker-width=640 ! eswstudentaction name=eswstudentactions0 config-file-path=/usr/lib/eswin/aarch64/ai-config/config_studentaction.txt ! eswcloseup1.sink_1`)
	// fmt.Println(strings.ReplaceAll(string(bytes), " ! ", "\\ \n   ! "))

	// mapt := map[string]string{}
	// mapt["t1"] = "1"
	// if _, ok := mapt["t1"]; ok {
	// 	fmt.Println("t1ok")
	// }
	// if _, ok := mapt["t2"]; ok {
	// 	fmt.Println("t2")
	// }
	var m, n int
	fmt.Scanln(&m, &n) //newline

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&res[i][j]) //newline=space
		}
	}
	fmt.Println(res)
}
