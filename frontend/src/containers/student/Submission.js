import React, { Component } from 'react';
import Highlight from 'react-highlight';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/student/Submission.css';
import "../../../node_modules/highlight.js/styles/atom-one-light.css";

/**
* Form where students can upload and submit a project. 
*/
class StudentSubmission extends Component {

  constructor(props) {
    super(props);
    let due = new Date("2016-11-07T19:06:00");
    this.state = {
      due: due,
      delta: due.getTime() - Date.now()
    }
  }

  componentWillMount () {
      const script = document.createElement("script");

      const scriptText = document.createTextNode(`
        function activateTab(event, id) {
          tablinks = document.getElementsByClassName("code-feedback");
          for (i = 0; i < tablinks.length; i++) {
              tablinks[i].className = tablinks[i].className.replace(" active", "");
          }
          event.currentTarget.className += " active";
        }
        `);

      script.appendChild(scriptText);
      document.head.appendChild(script);
  }

  componentDidMount() {
    this.timerID = setInterval(
      () => this.tick(),
      1000
    );
  }

  componentWillUnmount() {
    clearInterval(this.timerID);
  }

  tick() {
    this.setState({
      delta: this.state.due.getTime() - Date.now()
    });
  }

  back() {
    this.props.history.goBack();
  }

  activateTab(id) {
    for (var ref in this.refs) {
      this.refs[ref].className = "code-feedback";
      if (this.refs[ref].id === id) {
        this.refs[ref].className = "code-feedback active";
      }
    }
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path="Submission" />
          { this.state.delta > 0 ?
            <div className="text-center">
              <h1 className="blue text-center">Time to Deadline:</h1>
              <h1 className="dark-gray text-center">
                <span>{Math.floor((this.state.delta / (1000 * 60 * 60 * 24)) + (1 / 1440))}</span> days <span>{Math.floor((this.state.delta / (1000 * 60 * 60)) + (1 / 60)) % 24}</span> hours <span>{Math.ceil(this.state.delta / (1000 * 60)) % 60}</span> minutes
              </h1>
              <br />
              <h1 className="blue text-center">Add/Edit Submission</h1>
              <input id="upload" type="file" />
            </div>
            :
            <div>
              <h1 className="dark-gray text-center">Submission Feedback:</h1>
              <div id="left-feedback">
                <button onClick={() => this.activateTab("submission")}>Submission</button>
                <button onClick={() => this.activateTab("script")}>Test Script</button>
                <div id="submission" ref="code-feedback active" className="code-feedback active">
                  <Highlight className="cpp">{`
  #include <iostream>
  int main(int argc, char *argv[]) {
    for (auto i = 0; i <= 0xFFFF; i++)
      cout << "Hello, World!" << endl;
    return 1;
  }
                  `}</Highlight>
                </div>
                <div id="script" ref="code-feedback" className="code-feedback">
                  <Highlight className="cpp">{`
  #include <iostream>
  int main(int argc, char *argv[]) {
    cout << "test_case_3: Off-by-one error" << endl;
    return 1;
  }
                  `}</Highlight>
                </div>
              </div>
              <div id="right-feedback">
                <h2 id="feedback-score" className="gray">Score: 67%</h2>
                <br />
                <p className="error">test_case_3: Off-by-one error</p>
              </div>
            </div>
          }
        </div>
      </div>
    );
  }
}

export default StudentSubmission;
