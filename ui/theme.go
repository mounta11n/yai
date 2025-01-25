package ui

func GetTheme() []byte {
	return []byte(`
{
  "document": {
    "block_prefix": "\n",
    "block_suffix": "\n",
    "color": "#FFFFFF",
    "background_color": "#282C34",
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
    "color": "#FFAA00",
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
    "color": "#FFFFFF",
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
    "color": "#00D7FF",
    "underline": true
  },
  "link_text": {
    "color": "#00D7FF",
    "bold": true
  },
  "image": {
    "color": "#FFFFFF",
    "underline": true
  },
  "image_text": {
    "color": "#FFFFFF",
    "format": "Image: {{.text}} →"
  },
  "code": {
    "prefix": " ",
    "suffix": " ",
    "color": "#ffaa00",
    "background_color": "#262427"
  },
  "code_block": {
    "color": "#55557f",
    "margin": 2,
    "chroma": {
      "text": {
        "color": "#FFFFFF"
      },
      "error": {
        "color": "#FFFFFF",
        "background_color": "#ff557f"
      },
      "comment": {
        "color": "#8A8A8A"
      },
      "comment_preproc": {
        "color": "#8A8A8A"
      },
      "keyword": {
        "color": "#ff557f"
      },
      "keyword_reserved": {
        "color": "#ff557f"
      },
      "keyword_namespace": {
        "color": "#ff557f"
      },
      "keyword_type": {
        "color": "#aaffff"
      },
      "operator": {
        "color": "#FFFFFF"
      },
      "punctuation": {
        "color": "#FFFFFF"
      },
      "name": {
        "color": "#FFFFFF"
      },
      "name_builtin": {
        "color": "#FFFFFF"
      },
      "name_tag": {
        "color": "#55ffff"
      },
      "name_attribute": {
        "color": "#FFFFFF"
      },
      "name_class": {
        "color": "#55ffff"
      },
      "name_constant": {
        "color": "#ffaa7f"
      },
      "name_decorator": {
        "color": "#ff557f"
      },
      "name_exception": {},
      "name_function": {
        "color": "#aaff00"
      },
      "name_other": {},
      "literal": {},
      "literal_number": {
        "color": "#aaaaff"
      },
      "literal_date": {},
      "literal_string": {
        "color": "#ffaa7f"
      },
      "literal_string_escape": {
        "color": "#ffff7f"
      },
      "generic_deleted": {
        "color": "#ff557f"
      },
      "generic_emph": {
        "italic": true
      },
      "generic_inserted": {
        "color": "#55ffff"
      },
      "generic_strong": {
        "bold": true
      },
      "generic_subheading": {
        "color": "#767676"
      },
      "background": {
        "background_color": "#282C34"
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
