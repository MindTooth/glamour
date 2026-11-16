package styles

import "charm.land/glamour/v2/ansi"

// SrceryStyleConfig is the srcery style.
// Colors follow srcery-vim and the canonical palette in @srcery-colors/srcery-palette.
var SrceryStyleConfig = ansi.StyleConfig{
	Document: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			BlockPrefix: "\n",
			BlockSuffix: "\n",
			Color:       stringPtr("#FCE8C3"),
		},
		Margin: uintPtr(defaultMargin),
	},
	BlockQuote: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			Color: stringPtr("#2C78BF"),
		},
		Indent:      uintPtr(1),
		IndentToken: stringPtr("│ "),
	},
	Paragraph: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{},
	},
	List: ansi.StyleList{
		StyleBlock: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{},
		},
		LevelIndent: defaultListIndent,
	},
	Heading: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			BlockSuffix:     "\n",
			BackgroundColor: stringPtr("#1C1B19"),
			Bold:            boolPtr(true),
		},
	},
	H1: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			Prefix: "# ",
			Color:  stringPtr("#68A8E4"),
		},
	},
	H2: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			Prefix: defaultHeadingPrefixH2,
			Color:  stringPtr("#FBB829"),
		},
	},
	H3: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			Prefix: defaultHeadingPrefixH3,
			Color:  stringPtr("#FED06E"),
		},
	},
	H4: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			Prefix: defaultHeadingPrefixH4,
			Color:  stringPtr("#519F50"),
		},
	},
	H5: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			Prefix: defaultHeadingPrefixH5,
			Color:  stringPtr("#E02C6D"),
		},
	},
	H6: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			Prefix: defaultHeadingPrefixH6,
			Color:  stringPtr("#FF5C8F"),
		},
	},
	Text: ansi.StylePrimitive{},
	Strikethrough: ansi.StylePrimitive{
		CrossedOut: boolPtr(true),
	},
	Emph: ansi.StylePrimitive{
		Italic: boolPtr(true),
	},
	Strong: ansi.StylePrimitive{
		Bold: boolPtr(true),
	},
	HorizontalRule: ansi.StylePrimitive{
		Color:  stringPtr("#2C78BF"),
		Format: "\n────────\n",
	},
	Item: ansi.StylePrimitive{
		BlockPrefix: "• ",
	},
	Enumeration: ansi.StylePrimitive{
		BlockPrefix: ". ",
	},
	Task: ansi.StyleTask{
		StylePrimitive: ansi.StylePrimitive{},
		Ticked:         "[x] ",
		Unticked:       defaultTaskUntickedPrefix,
	},
	Link: ansi.StylePrimitive{
		Color: stringPtr("#98BC37"),
	},
	LinkText: ansi.StylePrimitive{
		Color:     stringPtr("#C5B088"),
		Underline: boolPtr(true),
	},
	Image: ansi.StylePrimitive{
		Color: stringPtr("#98BC37"),
	},
	ImageText: ansi.StylePrimitive{
		Color:     stringPtr("#C5B088"),
		Underline: boolPtr(true),
		Format:    defaultImageFormat,
	},
	Code: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			Color: stringPtr("#2C78BF"),
		},
	},
	CodeBlock: ansi.StyleCodeBlock{
		StyleBlock: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: stringPtr("#FCE8C3"),
			},
			Margin: uintPtr(defaultMargin),
		},
		Chroma: &ansi.Chroma{
			Text: ansi.StylePrimitive{
				Color: stringPtr("#FCE8C3"),
			},
			Error: ansi.StylePrimitive{
				Color:           stringPtr("#FCE8C3"),
				BackgroundColor: stringPtr("#EF2F27"),
				Bold:            boolPtr(true),
			},
			Comment: ansi.StylePrimitive{
				Color:  stringPtr("#917E6B"),
				Italic: boolPtr(true),
			},
			CommentPreproc: ansi.StylePrimitive{
				Color: stringPtr("#0AAEB3"),
			},
			Keyword: ansi.StylePrimitive{
				Color: stringPtr("#EF2F27"),
			},
			KeywordReserved: ansi.StylePrimitive{
				Color: stringPtr("#EF2F27"),
			},
			KeywordNamespace: ansi.StylePrimitive{
				Color: stringPtr("#F75341"),
			},
			KeywordType: ansi.StylePrimitive{
				Color:  stringPtr("#68A8E4"),
				Italic: boolPtr(true),
			},
			Operator: ansi.StylePrimitive{
				Color: stringPtr("#C5B088"),
			},
			Punctuation: ansi.StylePrimitive{
				Color: stringPtr("#917E6B"),
			},
			Name: ansi.StylePrimitive{
				Color: stringPtr("#FCE8C3"),
			},
			NameBuiltin: ansi.StylePrimitive{
				Color: stringPtr("#FBB829"),
			},
			NameTag: ansi.StylePrimitive{
				Color: stringPtr("#2C78BF"),
			},
			NameAttribute: ansi.StylePrimitive{
				Color: stringPtr("#68A8E4"),
			},
			NameClass: ansi.StylePrimitive{
				Color:  stringPtr("#68A8E4"),
				Italic: boolPtr(true),
			},
			NameConstant: ansi.StylePrimitive{
				Color: stringPtr("#FF5C8F"),
			},
			NameDecorator: ansi.StylePrimitive{
				Color: stringPtr("#0AAEB3"),
			},
			NameException: ansi.StylePrimitive{
				Color: stringPtr("#EF2F27"),
			},
			NameFunction: ansi.StylePrimitive{
				Color: stringPtr("#FBB829"),
			},
			NameOther: ansi.StylePrimitive{},
			Literal:   ansi.StylePrimitive{},
			LiteralNumber: ansi.StylePrimitive{
				Color: stringPtr("#FF5C8F"),
			},
			LiteralDate: ansi.StylePrimitive{
				Color: stringPtr("#FF5C8F"),
			},
			LiteralString: ansi.StylePrimitive{
				Color: stringPtr("#98BC37"),
			},
			LiteralStringEscape: ansi.StylePrimitive{
				Color: stringPtr("#FBB829"),
			},
			GenericDeleted: ansi.StylePrimitive{
				Color: stringPtr("#F75341"),
			},
			GenericEmph: ansi.StylePrimitive{
				Italic: boolPtr(true),
			},
			GenericInserted: ansi.StylePrimitive{
				Color: stringPtr("#98BC37"),
			},
			GenericStrong: ansi.StylePrimitive{
				Bold: boolPtr(true),
			},
			GenericSubheading: ansi.StylePrimitive{
				Color: stringPtr("#0AAEB3"),
			},
			Background: ansi.StylePrimitive{},
		},
	},
	Table: ansi.StyleTable{
		StyleBlock: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{},
		},
		CenterSeparator: stringPtr("┼"),
		ColumnSeparator: stringPtr("│"),
		RowSeparator:    stringPtr("─"),
	},
	DefinitionList: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{},
	},
	DefinitionTerm: ansi.StylePrimitive{},
	DefinitionDescription: ansi.StylePrimitive{
		BlockPrefix: defaultArrowBlockPrefix,
	},
	HTMLBlock: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{},
	},
	HTMLSpan: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{},
	},
}
