// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package examples

type exampleKey int

const (
	exampleHello exampleKey = iota
	exampleTimer
	exampleTodo
	exampleImmTodo
	exampleMarkdown
	exampleLatency
)

type _Imm_source struct {
	file string
	src  string
}

type _Imm_exampleSource map[exampleKey]*source

var sources = newExampleSource(func(es *exampleSource) {
	es.Set(exampleHello, new(source).setFile("hellomessage/hello_message.go"))
	es.Set(exampleTimer, new(source).setFile("timer/timer.go"))
	es.Set(exampleTodo, new(source).setFile("todoapp/todo_app.go"))
	es.Set(exampleImmTodo, new(source).setFile("immtodoapp/todo_app.go"))
	es.Set(exampleMarkdown, new(source).setFile("markdowneditor/markdown_editor.go"))
	es.Set(exampleLatency, new(source).setFile("sites/latency/latency.go"))
})

var fetchStarted bool

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

var latencyJsx = `n/a`
