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
                <button>Submission</button>
                <div id="code-feedback">
                  <Highlight className="cpp">{`
  #include <iostream>
  int main(int argc, char *argv[]) {
    for (auto i = 0; i < 0xFFFF; i++)
      cout << "Hello, World!" << endl;
    return 1;
  }
                  `}</Highlight>
                </div>
              </div>
              <div id="right-feedback">
                <p>jfkajsfaskfjakfj eafjafhakjnjkhkgahskjnask akjsnc jkash cjkascas</p>
              </div>
            </div>
          }
        </div>
      </div>
    );
  }
}

export default StudentSubmission;
