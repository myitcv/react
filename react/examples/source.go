package main

import (
	. "github.com/myitcv/gopherjs/react"
	. "github.com/myitcv/gopherjs/react/examples/hellomessage"
	. "github.com/myitcv/gopherjs/react/examples/markdowneditor"
	. "github.com/myitcv/gopherjs/react/examples/timer"
	. "github.com/myitcv/gopherjs/react/examples/todoapp"
)

type example struct {
	title     string
	jsxSource string
	goSource  string
	elem      func() Element
}

var examples = [...]example{
	example{
		"A Simple Example",
		helloMessageJsx,
		"hellomessage/hello_message.go",
		func() Element {
			return HelloMessage(
				HelloMessageProps{Name: "Jane"},
			)
		},
	},
	example{
		"A Stateful Component",
		timerJsx,
		"timer/timer.go",
		func() Element {
			return Timer()
		},
	},
	example{
		"An Application",
		applicationJsx,
		"todoapp/todo_app.go",
		func() Element {
			return TodoApp()
		},
	},
	example{
		"A Component Using External Plugins",
		markdownEditorJsx,
		"markdowneditor/markdown_editor.go",
		func() Element {
			return MarkdownEditor()
		},
	},
}

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
