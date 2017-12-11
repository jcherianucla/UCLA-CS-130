import React, { Component } from 'react';
import Highlight from 'react-highlight';
import { Bar, BarChart, XAxis, YAxis } from 'recharts';
import Header from '../shared/Header.js'
import SidePanel from '../shared/SidePanel.js'
import '../styles/Project.css';
import '../styles/circle.css';
import "../../node_modules/highlight.js/styles/atom-one-light.css";

var data = [
  {name: '0-10', students: 0},
  {name: '10-20', students: 0},
  {name: '20-30', students: 0},
  {name: '30-40', students: 0},
  {name: '40-50', students: 0},
  {name: '50-60', students: 0},
  {name: '60-70', students: 0},
  {name: '70-80', students: 0},
  {name: '80-90', students: 0},
  {name: '90-100', students: 0},
];

/**
* Form where students and professors can view details about a project
*/
class Project extends Component {

  constructor(props) {
    super(props);
    let due = new Date("2016-11-07T19:06:00");
    this.state = {
      due: due,
      delta: due.getTime() - Date.now(),
      class_name: '',
      project_name: '',
      key: Math.random()
    }
  }

  loadCurrentClass() {
    let token = localStorage.getItem('token');
    let self = this
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + self.props.match.params.class_id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + token,
        'Accept': 'application/json'
      },
    })
    .then((response) => response.json())
    .then(function (responseJSON) {
      console.log(responseJSON);
      if (responseJSON.class !== null && responseJSON.class.name !== null) {
        self.setState({'class_name': responseJSON.class.name});
      }
    });
  }

  loadCurrentProject() {
    let token = localStorage.getItem('token');
    let self = this
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + self.props.match.params.class_id + '/assignments/' + self.props.match.params.project_id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + token,
        'Accept': 'application/json'
      },
    })
    .then((response) => response.json())
    .then(function (responseJSON) {
      console.log(responseJSON);
      if (responseJSON.assignment && responseJSON.assignment.name) {
        self.setState({'project_name': responseJSON.assignment.name});
      }
      if (localStorage.getItem('role') === 'professor') {
        self.refs.circle.className += " p" + responseJSON.analytics.num_submissions;
        self.refs.percent.innerHTML = responseJSON.analytics.num_submissions + " %";
        var scores = [0, 21, 4, 22, 9];
        scores.forEach(function(element) {
          data[Math.floor(element / 10)]['students'] += 1;
        });
        self.refs.chart.data = data;
        self.setState({key: Math.random()});
      }
    });
  }

  componentWillMount () {
    if (localStorage.getItem('role') === "" || localStorage.getItem('token') === "") {
      this.props.history.push('/login');
    }
    this.loadCurrentClass();
    this.loadCurrentProject();
    let role = localStorage.getItem('role');
    if (role == null) {
      this.props.history.push('/login');
    }
    this.setState({role: role});

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
      if (ref === "submission-button" || ref === "submission-button active") {
        this.refs[ref].className = "submission-button";
        if (this.refs[ref].id === id) {
          this.refs[ref].className = "submission-button active";
        }
      } else {
        this.refs[ref].className = "code-feedback";
        if (this.refs[ref].id === id) {
          this.refs[ref].className = "code-feedback active";
        }
      }
    }
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path={["Classes", ["Projects", this.props.match.params.class_id], ["Submission", this.props.match.params.class_id, this.props.match.params.project_id]]} props={this.state}/>
          <div>
            <p className="dark-gray"><b>Project Description:</b> Project 2 Description goes here. It will be whatever the professor types in on the create for the project creation. We will update it to be something dynamic when we hook up the frontend and backend soon </p>
          </div>
          { this.state.role === "student" ?
            <div>
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
                    <button id="submission" ref="submission-button active" className="submission-button active" onClick={() => this.activateTab("submission")}>Submission</button>
                    <button id="script" ref="submission-button" className="submission-button" onClick={() => this.activateTab("script")}>Test Script</button>
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
                    <p className="red">test_case_3: Off-by-one error</p>
                  </div>
                </div>
              }
            </div>
          :
          <div>
            <h1 className="blue text-center">Total Submissions</h1>
            <div className="center-object">
              <div ref="circle" className="c100 big blue-circle">
                <span ref="percent"></span>
                <div className="slice">
                  <div className="bar"></div>
                  <div className="fill"></div>
                </div>
              </div>
            </div>
            <h1 className="blue text-center">Score Breakdown</h1>
            <div className="center-object">
              <BarChart ref="chart" width={800} height={300} data={data} key={this.state.key}>
                <XAxis dataKey="name" />
                <YAxis name="students" />
                <Bar type="monotone" dataKey="students" barSize={30} fill="#8884d8"/>
              </BarChart>
            </div>
            <br /><br />
          </div>
          }
        </div>
      </div>
    );
  }
}

export default Project;
