import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/student/Submission.css';

/**
* Form where students can upload and submit a project. 
*/
class StudentSubmission extends Component {

  constructor(props) {
    super(props);
    let due = new Date("2017-12-06T19:06:00");
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
          <h1 className="dark-gray text-center">Time to Deadline:</h1>
          { this.state.delta > 0 ? 
            <h1 className="dark-gray text-center">
              <span>{Math.floor((this.state.delta / (1000 * 60 * 60 * 24)) + (1 / 1440))}</span> days <span>{Math.floor((this.state.delta / (1000 * 60 * 60)) + (1 / 60)) % 24}</span> hours <span>{Math.ceil(this.state.delta / (1000 * 60)) % 60}</span> minutes
            </h1>
            :
            <h1 className="dark-gray text-center">
              The submission is now closed
            </h1>
          }
        </div>
      </div>
    );
  }
}

export default StudentSubmission;
