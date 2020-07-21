package i18n

// TranslationSet is a set of localised strings for a given language
type TranslationSet struct {
	NotEnoughSpace                             string
	ProjectTitle                               string
	MainTitle                                  string
	GlobalTitle                                string
	Navigate                                   string
	Menu                                       string
	Execute                                    string
	Scroll                                     string
	Close                                      string
	ErrorTitle                                 string
	RunningSubprocess                          string
	NoViewMachingNewLineFocusedSwitchStatement string
	OpenConfig                                 string
	EditConfig                                 string
	AnonymousReportingTitle                    string
	AnonymousReportingPrompt                   string
	ConfirmQuit                                string
	ErrorOccurred                              string
	ConnectionFailed                           string
	UnattachableContainerError                 string
	CannotAttachStoppedContainerError          string
	CannotAccessDockerSocketError              string
	CannotKillChildError                       string

	Donate                     string
	Cancel                     string
	CustomCommandTitle         string
	BulkCommandTitle           string
	Remove                     string
	HideStopped                string
	ForceRemove                string
	RemoveWithVolumes          string
	MustForceToRemoveContainer string
	Confirm                    string
	Return                     string
	FocusMain                  string
	StopContainer              string
	RestartingStatus           string
	StoppingStatus             string
	RemovingStatus             string
	RunningCustomCommandStatus string
	RunningBulkCommandStatus   string
	RemoveService              string
	Stop                       string
	Restart                    string
	Rebuild                    string
	Recreate                   string
	PreviousContext            string
	NextContext                string
	Attach                     string
	ViewLogs                   string
	ServicesTitle              string
	ContainersTitle            string
	StandaloneContainersTitle  string
	TopTitle                   string
	ImagesTitle                string
	VolumesTitle               string
	NoContainers               string
	NoContainer                string
	NoImages                   string
	NoVolumes                  string
	RemoveImage                string
	RemoveVolume               string
	RemoveWithoutPrune         string
	PruneImages                string
	PruneContainers            string
	PruneVolumes               string
	ConfirmPruneContainers     string
	ConfirmStopContainers      string
	ConfirmRemoveContainers    string
	ConfirmPruneImages         string
	ConfirmPruneVolumes        string
	PruningStatus              string
	StopService                string
	PressEnterToReturn         string
	StopAllContainers          string
	RemoveAllContainers        string
	ViewRestartOptions         string
	RunCustomCommand           string
	ViewBulkCommands           string

	LogsTitle                string
	ConfigTitle              string
	DockerComposeConfigTitle string
	StatsTitle               string
	CreditsTitle             string
	ContainerConfigTitle     string

	No  string
	Yes string
}

func koreanSet() TranslationSet {
	return TranslationSet{
		PruningStatus:              "pruning",
		RemovingStatus:             "removing",
		RestartingStatus:           "restarting",
		StoppingStatus:             "stopping",
		RunningCustomCommandStatus: "running custom command",
		RunningBulkCommandStatus:   "running bulk command",

		RunningSubprocess:                          "running subprocess",
		NoViewMachingNewLineFocusedSwitchStatement: "No view matching newLineFocused switch statement",

		ErrorOccurred:                     "An error occurred! Please create an issue at https://github.com/jesseduffield/lazydocker/issues",
		ConnectionFailed:                  "connection to docker client failed. You may need to restart the docker client",
		UnattachableContainerError:        "Container does not support attaching. You must either run the service with the '-it' flag or use `stdin_open: true, tty: true` in the docker-compose.yml file",
		CannotAttachStoppedContainerError: "You cannot attach to a stopped container, you need to start it first (which you can actually do with the 'r' key) (yes I'm too lazy to do this automatically for you) (pretty cool that I get to communicate one-on-one with you in the form of an error message though)",
		CannotAccessDockerSocketError:     "Can't access docker socket at: unix:///var/run/docker.sock\nRun lazydocker as root or read https://docs.docker.com/install/linux/linux-postinstall/",
		CannotKillChildError:              "Waited three seconds for child process to stop. There may be an orphan process that continues to run on your system.",

		Donate:  "기부",
		Confirm: "확인",

		Return:              "return",
		FocusMain:           "focus main panel",
		Navigate:            "navigate",
		Execute:             "execute",
		Close:               "닫기",
		Menu:                "메뉴",
		Scroll:              "스크롤",
		OpenConfig:          "lazydocker 설정 열기",
		EditConfig:          "lazydocker 설정 편집",
		Cancel:              "취소",
		Remove:              "제거",
		HideStopped:         "Hide/Show stopped containers",
		ForceRemove:         "강제 제거",
		RemoveWithVolumes:   "볼륨과 함께 제거",
		RemoveService:       "컨테이너 제거",
		Stop:                "정지",
		Restart:             "재실행",
		Rebuild:             "재빌드",
		Recreate:            "재생성",
		PreviousContext:     "이전 탭",
		NextContext:         "다음 탭",
		Attach:              "연결",
		ViewLogs:            "로그 보기",
		RemoveImage:         "이미지 제거",
		RemoveVolume:        "볼륨 제거",
		RemoveWithoutPrune:  "remove without deleting untagged parents",
		PruneContainers:     "정지한 컨테이너 청소",
		PruneVolumes:        "미사용 볼륨 청소",
		PruneImages:         "미사용 이미지 청소",
		StopAllContainers:   "모든 컨테이너 정지",
		RemoveAllContainers: "모든 컨테이너 정지 (강제)",
		ViewRestartOptions:  "모든 재실행 옵션 보기",
		RunCustomCommand:    "미리 정의한 사용자 명령어 실행",
		ViewBulkCommands:    "벌크 명령어 보기",

		AnonymousReportingTitle:  "lazydocker가 더 좋아지도록 도와주세요",
		AnonymousReportingPrompt: "익명으로 데이터를 제출해주시면 lazydocker를 더 좋게 만들 수 있습니다. 도와주시겠어요?",

		GlobalTitle:               "Global",
		MainTitle:                 "Main",
		ProjectTitle:              "프로젝트",
		ServicesTitle:             "서비스",
		ContainersTitle:           "컨테이너",
		StandaloneContainersTitle: "컨테이너 - 단독 실행",
		ImagesTitle:               "이미지",
		VolumesTitle:              "볼륨",
		CustomCommandTitle:        "사용자 명령어:",
		BulkCommandTitle:          "벌크 명렁어:",
		ErrorTitle:                "Error",
		LogsTitle:                 "Logs",
		ConfigTitle:               "Config",
		DockerComposeConfigTitle:  "Docker-Compose Config",
		TopTitle:                  "Top",
		StatsTitle:                "Stats",
		CreditsTitle:              "About",
		ContainerConfigTitle:      "Container Config",

		NoContainers: "컨테이너 없음",
		NoContainer:  "컨테이너 없음",
		NoImages:     "이미지 없음",
		NoVolumes:    "볼륨 없음",

		ConfirmQuit:                "이 화면을 종료할까요?",
		MustForceToRemoveContainer: "실행 중인 컨테이너는 제거할 수 없습니다. 강제로 제거할까요?",
		NotEnoughSpace:             "정보 영역을 출력하려면 더 넓은 화면이 필요합니다",
		ConfirmPruneImages:         "사용하지 않는 이미지를 모두 청소할까요?",
		ConfirmPruneContainers:     "정지한 컨테이너를 모두 청소할까요?",
		ConfirmStopContainers:      "컨테이너를 모두 정지할까요?",
		ConfirmRemoveContainers:    "모든 컨테이너를 제거할까요?",
		ConfirmPruneVolumes:        "사용하지 않는 볼륨을 모두 청소할까요?",
		StopService:                "이 서비스의 컨테이너를 정지할까요?",
		StopContainer:              "이 컨테이너를 정지할까요?",
		PressEnterToReturn:         "Enter 를 입력하면 lazydocker 로 돌아갑니다. " +
									"(이 메시지는 설정에 `gui.returnImmediately: true` 를 반영하면 나오지 않습니다.)",

		No:  "아니요",
		Yes: "예",
	}
}
