import React from "react";
import ReactDOM from "react-dom";
import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import App from "./App";
import TestQuestionAndAnswer from "./TestQuestionAndAnswer";

ReactDOM.render(
  <Router>
    <Routes>
      <Route path="/" element={<App />} />
      <Route
        path="/text-question-and-answer"
        element={<TestQuestionAndAnswer />}
      />
    </Routes>
  </Router>,

  document.getElementById("root")
);
