const question = [
  {
    que: "Which Company Developed Android",
    a: "Google",
    b: "Android",
    c: "Android inc.",
    d: "Nokia",
    correct: "b",
  },
  {
    que: "Android is ... ",
    a: "An operating system ",
    b: " A web browser",
    c: " a web server",
    d: " None of above",
    correct: "a",
  },
  {
    que: " Which of the following virtual machine is used by Android operating system",
    a: " JVM",
    b: " Dalvik virtual machine",
    c: " Simple virtual machine",
    d: " None of above",
    correct: "b",
  },
  {
    que: "Android is based on which of following language ? ",
    a: "Java ",
    b: "C++ ",
    c: "C ",
    d: "kotlin",
    correct: "a",
  },
  {
    que: " APK stands for ?",
    a: " Android Phone kit",
    b: " Android Page kit",
    c: " Android package kit",
    d: " Application page kit",
    correct: "c",
  },
  {
    que: "What does API stand for ?",
    a: " Application programming interface",
    b: "Android programming interface ",
    c: " Android page interface",
    d: " Application page interface",
    correct: "a",
  },
  {
    que: " How to stop the services in android?",
    a: " By using StopSelf() and stopService() method",
    b: " By using finish() method",
    c: " By using system.exit() method",
    d: " exit()",
    correct: "a",
  },
  {
    que: "Which of the following kernel is used in Android? ",
    a: " MAC",
    b: "Windows ",
    c: "Linux ",
    d: "Redhat ",
    correct: "c",
  },
  {
    que: "Does Android support other languages than java? ",
    a: " Yes",
    b: " No",
    c: " May be",
    d: " can't say",
    correct: "a",
  },
  {
    que: "Wwhich of the following method is used to handle what happens after clicking a button ",
    a: " onClick",
    b: " onCreate",
    c: " onSelect",
    d: " None of above",
    correct: "a",
  },
  {
    que: "Which of the following is contained in the src folder? ",
    a: " XML ",
    b: " Java source code",
    c: " Manifest",
    d: " JSON ",
    correct: "b",
  },
  {
    que: "Which of the above is parent class of Activity ? ",
    a: " context",
    b: " object",
    c: " ContextThemeWrapper",
    d: " None of above ",
    correct: "c",
  },
  {
    que: " which of following method is used for lo debug in android ?",

    a: " Log.r() ",
    b: " Log.R()",
    c: " Log.d()",
    d: " Log.D()",
    correct: "c",
  },
  {
    que: " All layout classes are the subclasses of .. ?",
    a: " android.view.View",
    b: " android.view.ViewGroup",
    c: " android.widget",
    d: " None of above",
    correct: "b",
  },

  {
    que: "Android is licensed under which open source licesing license ? ",
    a: " Gnus GPL",
    b: " Apache/MIT",
    c: " OSS",
    d: " Sourceforge",
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
    <div class='btnclass'>
    <a href = 'androiddev.html'> <button id="res" >Exit </button> </a>
    <a href = 'cer.html'> <button id="res">Get Certificate </button> </a>
</div>
    `;
};
// call
loadquestion();

