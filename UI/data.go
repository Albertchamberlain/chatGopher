package UI

import "fyne.io/fyne/v2"

// Tutorial defines the data structure for a tutorial
type AiModel struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// Tutorials defines the metadata for each tutorial
	AiModels = map[string]AiModel{
		"Welcome": {"Welcome", "", welcomeScreen, true},
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
