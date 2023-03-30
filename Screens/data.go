package UI

import "fyne.io/fyne/v2"

// AiModel defines the model's screen
type AiModel struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// AiModels defines the metadata for each model's screen
	AiModels = map[string]AiModel{
		"Welcome": {"Welcome",
			"",
			welcomeScreen,
			true,
		},
		"ChatGPT": {"ChatGPT",
			"ChatGPT Model(dafault GPT3.5)",
			chatGPTScreen,
			true,
		},
		"GPT-3": {"GPT-3",
			"GPT-3 Model",
			gpt3Screen,
			true,
		},
		"GPT-4": {"GPT-4",
			"GPT-4 Model",
			gpt4Screen,
			true,
		},
		"DALL·E2": {"DALL·E2",
			"DALL-E 2 image generation",
			dalle2Screen,
			true,
		},
		"Whisper": {"Whisper",
			"Whisper Model Audio Speech-To-Text",
			whisperScreen,
			true,
		},
	}

	// AiModelsIndex  defines how our AiModels should be laid out in the index tree
	AiModelsIndex = map[string][]string{
		"": {"Welcome", "ChatGPT", "GPT-3", "GPT-4", "DALL·E2", "Whisper"},
	}
)
