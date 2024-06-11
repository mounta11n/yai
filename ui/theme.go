package ui

func GetTheme() []byte {
	return []byte(`
{
  "document": {
    "block_prefix": "\n",
    "block_suffix": "\n",
    "color": "#FCFCFA",
    "background_color": "#262427",
    "margin": 2
  },
  "block_quote": {
    "indent": 1,
    "indent_token": "│ "
  },
  "paragraph": {},
  "list": {
    "level_indent": 2
  },
  "heading": {
    "block_suffix": "\n",
    "color": "#FC9867",
    "bold": true
  },
  "h1": {
    "prefix": " ",
    "suffix": " ",
    "bold": true
  },
  "h2": {
    "prefix": "## "
  },
  "h3": {
    "prefix": "### "
  },
  "h4": {
    "prefix": "#### "
  },
  "h5": {
    "prefix": "##### "
  },
  "h6": {
    "prefix": "###### ",
    "bold": false
  },
  "text": {},
  "strikethrough": {
    "crossed_out": true
  },
  "emph": {
    "italic": true
  },
  "strong": {
    "bold": true
  },
  "hr": {
    "color": "#FCFCFA",
    "format": "\n--------\n"
  },
  "item": {
    "block_prefix": "• "
  },
  "enumeration": {
    "block_prefix": ". "
  },
  "task": {
    "ticked": "[✓] ",
    "unticked": "[ ] "
  },
  "link": {
    "color": "#AEE8F4",
    "underline": true
  },
  "link_text": {
    "color": "#AEE8F4",
    "bold": true
  },
  "image": {
    "color": "#FCFCFA",
    "underline": true
  },
  "image_text": {
    "color": "#FCFCFA",
    "format": "Image: {{.text}} →"
  },
  "code": {
    "prefix": " ",
    "suffix": " ",
    "color": "#FFCA58",
    "background_color": "#262427"
  },
  "code_block": {
    "color": "#262427",
    "margin": 2,
    "chroma": {
      "text": {
        "color": "#FCFCFA"
      },
      "error": {
        "color": "#FCFCFA",
        "background_color": "#FF7272"
      },
      "comment": {
        "color": "#8B8B8A"
      },
      "comment_preproc": {
        "color": "#8B8B8A"
      },
      "keyword": {
        "color": "#FF7272"
      },
      "keyword_reserved": {
        "color": "#FF7272"
      },
      "keyword_namespace": {
        "color": "#FF7272"
      },
      "keyword_type": {
        "color": "#AEE8F4"
      },
      "operator": {
        "color": "#FCFCFA"
      },
      "punctuation": {
        "color": "#FCFCFA"
      },
      "name": {
        "color": "#FCFCFA"
      },
      "name_builtin": {
        "color": "#FCFCFA"
      },
      "name_tag": {
        "color": "#49CAE4"
      },
      "name_attribute": {
        "color": "#FCFCFA"
      },
      "name_class": {
        "color": "#49CAE4"
      },
      "name_constant": {
        "color": "#FFCA58"
      },
      "name_decorator": {
        "color": "#FF7272"
      },
      "name_exception": {},
      "name_function": {
        "color": "#BCDF59"
      },
      "name_other": {},
      "literal": {},
      "literal_number": {
        "color": "#A093E2"
      },
      "literal_date": {},
      "literal_string": {
        "color": "#FFCA58"
      },
      "literal_string_escape": {
        "color": "#FFCA58"
      },
      "generic_deleted": {
        "color": "#FF7272"
      },
      "generic_emph": {
        "italic": true
      },
      "generic_inserted": {
        "color": "#49CAE4"
      },
      "generic_strong": {
        "bold": true
      },
      "generic_subheading": {
        "color": "#777777"
      },
      "background": {
        "background_color": "#262427"
      }
    }
  },
  "table": {
    "center_separator": "┼",
    "column_separator": "│",
    "row_separator": "─"
  },
  "definition_list": {},
  "definition_term": {},
  "definition_description": {
    "block_prefix": "\n🠶 "
  },
  "html_block": {},
  "html_span": {}
}
		`)
}
