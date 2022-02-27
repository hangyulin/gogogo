import * as React from "react";
import "./App.css";
import { Container } from "@mui/material";
import Button from "@mui/material/Button";

function App() {
  return (
    <div>
      <Container>
        <Button variant="text" href="/text-question-and-answer">
          Text Question and Answer Feature
        </Button>
      </Container>
    </div>
  );
}

export default App;
