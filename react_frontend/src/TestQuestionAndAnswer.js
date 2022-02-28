import * as React from "react";
import "./App.css";
import { Container } from "@mui/material";
import TestQuestionAndAnswerBox from "./components/TextQuestionAndAnswerBox";

function TestQuestionAndAnswer() {
  return (
    <div>
      <Container>
        <TestQuestionAndAnswerBox />
      </Container>
    </div>
  );
}

export default TestQuestionAndAnswer;
