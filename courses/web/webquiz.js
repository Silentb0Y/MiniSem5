const question = [
  {
    que: "There are ___ levels of heading in HTML",
    a: "Three",
    b: "Four",
    c: "Five",
    d: "six",
    correct: "d",
  },
  {
    que: "Various INPUT types are",
    a: "WordPad",
    b: "WordPad, Notepad",
    c: "WordPad, Write pad, Notepad",
    d: "None of the above",
    correct: "c",
  },
  {
    que: " HTML is the method where ordinary text can be converted into",
    a: "ASCII Text",
    b: "EBCDIC Text",
    c: "EBCDIC Text",
    d: "EBCDIC Text",
    correct: "c",
  },
  {
    que: "Which HTML tag is used to define an internal style sheet? ",
    a: "<style> ",
    b: " <script>",
    c: " <css>",
    d: " <link>",
    correct: "b",
  },
  {
    que: "In HTML ___ tag contains the information about the current document such as title etc. ",
    a: " body",
    b: "TD ",
    c: "Head ",
    d: " None of above",
    correct: "c",
  },
  {
    que: " The ___ element can be used to identify your HTML file to the outside world",
    a: "Title",
    b: "Body ",
    c: "Head ",
    d: "None of the above ",
    correct: "a",
  },
  {
    que: "  The tag to give visual division between sections of the page and which causes the browser to draw an embossed line is",
    a: " <HL>",
    b: " <HR>",
    c: "<UR> ",
    d: " None of the above",
    correct: "b",
  },
  {
    que: " Which one of the following tags are used to insert graphics on the web page?",
    a: "<IMAGE> ",
    b: "<IMAGES> ",
    c: "<IMG> ",
    d: " <GRAPHICS",
    correct: "c",
  },
  {
    que: " You can redirect the webpage in JavaScript by using ………………… method.",
    a: "window. reload ",
    b: "window.location ",
    c: " page.location",
    d: " url.newlocation",
    correct: "b",
  },
  {
    que: " ………………….. is a built-in JavaScript function which can be used to execute another function after a given time interval. ",
    a: " Timeout( )",
    b: " Timeout( )",
    c: "setTimeout ( ) ",
    d: "  All of the above",
    correct: "c",
  },
  {
    que: " What does CSS stand for?",
    a: " Computer style sheets",
    b: " Creative style sheets",
    c: " Colorful style sheets",
    d: " cascading style sheets",
    correct: "d",
  },
  {
    que: "What does CSS stand for? ",
    a: " <stylsheet>mystyle.css</stylesheet>",
    b: " <link rel='stylesheet' type='text/css' href='mystyle.css'",
    c: "<style src ='mystyle.css' ",
    d: " None of the above",
    correct: "b",
  },
];
let index = 0;
let total = question.length;
let right = 0,
  wroung = 0;
const ques = document.getElementById("ques");
const optionInputs = document.querySelectorAll(".option");
const loadquestion = () => {
  if (index === total) {
    return end();
  }
  reset();
  const data = question[index];
  ques.innerText = `${index + 1} ${data.que}`;
  optionInputs[0].nextElementSibling.innerText = data.a;
  optionInputs[1].nextElementSibling.innerText = data.b;
  optionInputs[2].nextElementSibling.innerText = data.c;
  optionInputs[3].nextElementSibling.innerText = data.d;
};
const qui = () => {
  const ans = getAnswer();
  const data = question[index];
  if (ans == data.correct) {
    right++;
  } else {
    wroung++;
  }
  index++;
  loadquestion();
  return;
};
const getAnswer = () => {
  let answer;
  optionInputs.forEach((input) => {
    if (input.checked) {
      answer = input.value;
    }
  });
  return answer;
};

const reset = () => {
  optionInputs.forEach((input) => {
    input.checked = false;
  });
};

let username = localStorage.getItem("nameC");
document.getElementById("User").innerHTML = username;
const end = () => {
  document.getElementById("bx").innerHTML = `
  <h2 class="user">${username}</h2>
    <h2> ${right} / ${total} are correct</h2>
    <h3> Quiz is Ended </h3>
    <a href = 'androiddev.html'> <button id="res" class='btn'>Exit </button> </a>
    <a href = 'cer.html'> <button id="res" class='btn'>Get Certificate </button> </a>

    `;
};

// call
loadquestion();
