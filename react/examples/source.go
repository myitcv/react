// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package examples

import (
	r "github.com/myitcv/gopherjs/react"
	"github.com/myitcv/gopherjs/react/examples/hellomessage"
	"github.com/myitcv/gopherjs/react/examples/immtodoapp"
	"github.com/myitcv/gopherjs/react/examples/markdowneditor"
	"github.com/myitcv/gopherjs/react/examples/timer"
	"github.com/myitcv/gopherjs/react/examples/todoapp"
)

type _Imm_example struct {
	title        string
	message      string
	jsxSourceStr string
	goSourceFile string
	goSourceStr  string
	elem         func() r.Element
}

func newExample(title, message, jsxSourceStr, goSourceFile string, elem func() r.Element) *example {
	return new(example).WithMutable(func(e *example) {
		e.setTitle(title)
		e.setMessage(message)
		e.setJsxSourceStr(jsxSourceStr)
		e.setGoSourceFile(goSourceFile)
		e.setElem(elem)
	})
}

var fetchStarted bool

var examples = newExampleS([]*example{
	newExample(
		"A Simple Example",
		"The hellomessage.HelloMessage component demonstrates the simple use of a Props type.",
		helloMessageJsx,
		"hellomessage/hello_message.go",
		func() r.Element {
			return hellomessage.HelloMessage(
				hellomessage.HelloMessageProps{Name: "Jane"},
			)
		},
	),
	newExample(
		"A Stateful Component",
		"The timer.Timer component demonstrates the use of a State type.",
		timerJsx,
		"timer/timer.go",
		func() r.Element {
			return timer.Timer()
		},
	),
	newExample(
		"An Application",
		"The todoapp.TodoApp component demonstrates the use of state and event handling, but also the "+
			"problems of having a non-comparable state struct type.",
		applicationJsx,
		"todoapp/todo_app.go",
		func() r.Element {
			return todoapp.TodoApp()
		},
	),
	newExample(
		"An Application using github.com/myitcv/immutable",
		"The immtodoapp.TodoApp component is a reimplementation of todoapp.TodoApp using immutable data structures.",
		"n/a",
		"immtodoapp/todo_app.go",
		func() r.Element {
			return immtodoapp.TodoApp()
		},
	),
	newExample(
		"A Component Using External Plugins",
		"The markdowneditor.MarkdownEditor component demonstrates the use of an external Javascript library.",
		markdownEditorJsx,
		"markdowneditor/markdown_editor.go",
		func() r.Element {
			return markdowneditor.MarkdownEditor()
		},
	),
}...)

var helloMessageJsx = `class HelloMessage extends React.Component {
  render() {
    return <div>Hello {this.props.name}</div>;
  }
}`

var timerJsx = `class Timer extends React.Component {
  constructor(props) {
    super(props);
    this.state = {secondsElapsed: 0};
  }

  tick() {
    this.setState((prevState) => ({
      secondsElapsed: prevState.secondsElapsed + 1
    }));
  }

  componentDidMount() {
    this.interval = setInterval(() => this.tick(), 1000);
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  render() {
    return (
      <div>Seconds Elapsed: {this.state.secondsElapsed}</div>
    );
  }
}`

var applicationJsx = `class TodoApp extends React.Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.state = {items: [], text: ''};
  }

  render() {
    return (
      <div>
        <h3>TODO</h3>
        <TodoList items={this.state.items} />
        <form onSubmit={this.handleSubmit}>
          <input onChange={this.handleChange} value={this.state.text} />
          <button>{'Add #' + (this.state.items.length + 1)}</button>
        </form>
      </div>
    );
  }

  handleChange(e) {
    this.setState({text: e.target.value});
  }

  handleSubmit(e) {
    e.preventDefault();
    var newItem = {
      text: this.state.text,
      id: Date.now()
    };
    this.setState((prevState) => ({
      items: prevState.items.concat(newItem),
      text: ''
    }));
  }
}

class TodoList extends React.Component {
  render() {
    return (
      <ul>
        {this.props.items.map(item => (
          <li key={item.id}>{item.text}</li>
        ))}
      </ul>
    );
  }
}`

var markdownEditorJsx = `class MarkdownEditor extends React.Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
    this.state = {value: 'Type some *markdown* here!'};
  }

  handleChange() {
    this.setState({value: this.refs.textarea.value});
  }

  getRawMarkup() {
    var md = new Remarkable();
    return { __html: md.render(this.state.value) };
  }

  render() {
    return (
      <div className="MarkdownEditor">
        <h3>Input</h3>
        <textarea
          onChange={this.handleChange}
          ref="textarea"
          defaultValue={this.state.value} />
        <h3>Output</h3>
        <div
          className="content"
          dangerouslySetInnerHTML={this.getRawMarkup()}
        />
      </div>
    );
  }
}`
