import React, { useState } from "react";
import axios from "axios";
import { Grid, Paper, Avatar, TextField, Button } from "@mui/material";
import { LoadingButton } from "@mui/lab";
import MenuBookOutlinedIcon from "@mui/icons-material/MenuBookOutlined";

const TestQuestionAndAnswerBox = () => {
  const [context, setContext] = useState("");
  const [question, setQuestion] = useState("");
  const [answer, setAnswer] = useState("");
  const [loading, setLoading] = useState(false);

  const paperStyle = { padding: 40, height: "80vh", margin: "20px auto" };
  const avatarStyle = { backgroundColor: "blue" };
  const textStyle = { marginTop: "20px" };
  const buttonStyle = { marginTop: "30px", marginBottom: "30px" };

  function Submit() {
    if (context && question) {
      setLoading(true);
      let items = {
        context: context,
        question: question,
      };
      axios.post("/api/qna", items).then((res) => {
        setAnswer(res.data);
        setLoading(false);
      });
    }
  }
  return (
    <Grid>
      <Paper elevation={10} style={paperStyle}>
        <Grid align="center">
          <Avatar style={avatarStyle}>
            <MenuBookOutlinedIcon></MenuBookOutlinedIcon>
          </Avatar>
          <h2>Text Question and Answer</h2>
        </Grid>
        <TextField
          label="Text to Interpret"
          placeholder="e.g. Ontario is reporting 1,003 people hospitalized with COVID-19 on Friday, marking the lowest number the province has seen since last December — the beginning of the Omicron wave. The number of people in intensive care units also dropped, with the province reporting 297 people in ICUs. That marks the first time that figure has been under 300 since Jan. 5. The province is set to lift its proof of vaccination requirements next week as part of Ontario's reopening plan due to positive trends in public health indicators but the province's top doctor said mask requirements will remain in place for the time being. Chief Medical Officer of Health Dr. Kieran Moore said at a Thursday news briefing that masks remain 'an important tool in our tool box' when it comes to reducing transmission of the virus. Moore said given positive trends in public health indicators and the province's high vaccination rate, health officials are actively reviewing all directives to health-care providers, and hope to provide an update in the coming weeks. The number of hospitalizations reported Friday is down from 1,066 the day before and from 1,281 at the same time last week."
          fullWidth
          multiline
          rows={5}
          style={textStyle}
          onChange={(e) => setContext(e.target.value)}
          value={context}
          inputProps={{ maxLength: 3000 }}
        ></TextField>
        <TextField
          label="Question About the Text"
          placeholder="e.g. Why is proof of vaccination requirements lifted in Ontario?"
          fullWidth
          multiline
          rows={1}
          style={textStyle}
          onChange={(e) => setQuestion(e.target.value)}
          value={question}
          inputProps={{ maxLength: 200 }}
        ></TextField>

        <LoadingButton
          loading={loading}
          variant="contained"
          color="primary"
          fullWidth
          size="large"
          style={buttonStyle}
          onClick={() => {
            Submit();
          }}
        >
          Get Answer to Question
        </LoadingButton>

        <TextField
          disabled
          id="outlined-disabled"
          label="Answer to Question"
          fullWidth
          multiline
          rows={3}
          style={textStyle}
          value={answer}
        />
      </Paper>
    </Grid>
  );
};

export default TestQuestionAndAnswerBox;
