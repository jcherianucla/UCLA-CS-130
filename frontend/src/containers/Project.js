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
    let due = new Date();
    this.state = {
      due: due,
      delta: due.getTime() - Date.now(),
      class_name: '',
      project_name: '',
      key: Math.random(),
      loaded: '',
      submission_message: ''
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
        self.refs.description.innerHTML = responseJSON.assignment.name + ": " + responseJSON.assignment.description;
      }
      if (localStorage.getItem('role') === 'professor') {
        self.setState({loaded: true});
        self.refs.circle.className += " p" + Math.floor(responseJSON.analytics.num_submissions * 100);
        self.refs.percent.innerHTML = Math.floor(responseJSON.analytics.num_submissions * 100) + " %";
        if (responseJSON.analytics.score) {
          responseJSON.analytics.score.forEach(function(element) {
            data[Math.floor(element / 10)]['students'] += 1;
          });
        }
        self.refs.chart.data = data;
        self.setState({key: Math.random()});
      } else if (!responseJSON.submission) {
        self.setState({due: new Date(new Date(responseJSON.assignment.deadline).getTime() + 8 * 3600 * 1000)});
        self.tick();
        self.setState({loaded: true});
      } else {
        self.setState({loaded: true});
        console.log(self.state.due);
        console.log(responseJSON.assignment.deadline);
        self.refs.score.innerHTML = "Score: " + responseJSON.submission.submission.score;
        var feedback = "";
        responseJSON.submission.submission.post_results.map(function(item, key) {
          feedback += item + "\n";
        });
        self.refs.feedback.innerHTML = feedback;
        self.refs.grading.innerHTML = responseJSON.submission.grade_script;
        self.refs.submission.innerHTML = atob(responseJSON.submission.submission.file);
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
    let delta = this.state.due.getTime() - Date.now();
    if (delta <= 0 && this.state.delta > 0) {
      this.props.history.push('/classes/' + this.props.match.params.class_id + '/assignments/' + this.props.match.params.project_id);
    } else {
      this.setState({
        delta: delta
      });
    }
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
      } else if (ref === "code-feedback" || ref === "code-feedback active") {
        this.refs[ref].className = "code-feedback";
        if (this.refs[ref].id === id) {
          this.refs[ref].className = "code-feedback active";
        }
      }
    }
  }


  submit() {
    let token = localStorage.getItem('token');
    this.setState({submission_message: "Last Submitted at " + new Date().toLocaleTimeString()});
    var data = new FormData();
    data.append('upload', this.refs.upload.files[0]);

    let self = this
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + self.props.match.params.class_id + '/assignments/' + self.props.match.params.project_id + '/submissions' , {
      method: 'POST',
      headers: {
        'Authorization': 'Bearer ' + token
      },
      body: data
    })
    .then((response) => response.json())
    .then(function (responseJSON) {
      console.log(responseJSON);
      if (responseJSON.message !== 'Success') {
        fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + self.props.match.params.class_id + '/assignments/' + self.props.match.params.project_id + '/submissions' , {
          method: 'PUT',
          headers: {
            'Authorization': 'Bearer ' + token
          },
          body: data
        })
        .then((response) => response.json())
        .then(function (responseJSON) {
          console.log(responseJSON);
          if (responseJSON.message !== 'Success') {
            self.refs.error.innerHTML = responseJSON.message;
          }
        });
      }
    });

  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path={["Classes", ["Projects", this.props.match.params.class_id], ["Submission", this.props.match.params.class_id, this.props.match.params.project_id]]} props={this.state}/>
          <p ref="error" className="red"></p>
          <div>
            <p ref="description" className="dark-gray"></p>
          </div>
            { this.state.loaded === true ?
            <div>
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
                    <div className="upload-btn-wrapper">
                      <input ref="upload" id="upload" className="btn" type="file" onChange={() => this.submit()} accept=".cpp" required/>
                      <button className="btn">Upload</button>
                      <label className="filename" id="filename">{this.state.submission_message}</label>
                    </div>
                  </div>
                  :
                  <div>
                    <h1 className="dark-gray text-center">Submission Feedback:</h1>
                    <div id="left-feedback">
                      <button id="submission" ref="submission-button active" className="submission-button active" onClick={() => this.activateTab("submission")}>Submission</button>
                      <button id="script" ref="submission-button" className="submission-button" onClick={() => this.activateTab("script")}>Test Script</button>
                      <div id="submission" ref="code-feedback active" className="code-feedback active">
                        <pre>
                          <code ref="submission" className="cpp hljs"></code>
                        </pre>
                      </div>
                      <div id="script" ref="code-feedback" className="code-feedback">
                        <pre>
                          <code ref="grading" className="cpp hljs"></code>
                        </pre>
                      </div>
                    </div>
                    <div id="right-feedback">
                      <h2 ref="score" id="feedback-score" className="gray"></h2>
                      <br />
                      <p ref="feedback" className="red error"></p>
                    </div>
                  </div>
                }
              <br /><br />
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
                  <XAxis dataKey="name"/>
                  <YAxis name="students" allowDecimals={false}/>
                  <Bar type="monotone" dataKey="students" barSize={30} fill="#8884d8"/>
                </BarChart>
              </div>
              <br /><br />
            </div>
            }
            </div>
            :
            <div></div>
          }
        </div>
      </div>
    );
  }
}

export default Project;
