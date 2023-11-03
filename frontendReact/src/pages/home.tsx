import "../home.css"
// import "../all.css"
import { Nav } from "../components/nav"

export function Home() {
    return (
        <div>
            {/* <Nav /> */}
            <div className="container"> 
                <div className="training-item">
                    <img
                        src="./images/web-development-png-web-development-1100.png"
                        alt=""
                    />
                    <h3>Full stack developer</h3>
                    <p>
                        Welcome to web-development Course In this course we are providing all
                        required resource to You <small>Click Now</small>
                    </p>
                    <a href="./courses/web/Web Development.html" target="blank"
                    ><button>Start</button>
                    </a>
                </div>
                <div className="training-item">
                    <img src="./images/android.png" alt="" />
                    <h3>Android Developer</h3>
                    <p>
                        Welcome to Android development Course In this course we are providing
                        all required resource to You <small>Click Now</small>
                    </p>
                    <a href="./courses/android/androiddev.html" target="blank"
                    ><button>Start</button>
                    </a>
                </div>
                <div className="training-item">
                    <img src="./images/cloud-data.png" alt="" />
                    <h3>Machine Learning</h3>
                    <p>
                        Welcome to Machine Learing Course In this course we are providing all
                        required resource to You <small>Click Now</small>
                    </p>
                    <a href="./courses/Machine/machine.html" target="_blank"
                    ><button>Start</button>
                    </a>
                </div>
                <div className="training-item">
                    <img src="./images/hacking.png" alt="" />
                    <h3>Cyber Security</h3>
                    <p>
                        Welcome to Cyber Security Course In this course we are providing all
                        required resource to You <small>Click Now</small>
                    </p>
                    <a href="./courses/CyberSecurity/Cybercon.html"
                    ><button>Start</button>
                    </a>
                </div>
                <div className="training-item">
                    <img src="./images/ds.jpg" alt="" />
                    <h3>Data Science</h3>
                    <p>
                        Welcome to Data Science Course In this course we are providing all
                        required resource to You <small>Click Now</small>
                    </p>
                    <a href="./courses/ds/dscon.html"><button>read more</button> </a>
                </div>
                <div className="training-item">
                    <img src="images/dsa.png" alt="" />
                    <h3 style={{ overflow: "hidden" }}>Data Structures & Algorithms</h3>
                    <p>
                        Welcome to Data Structures and Algorithm Course In this course we are
                        providing all required resource to You <small>Click Now</small>
                    </p>
                    <a href="./courses/DSA/dsa.html"><button>read more</button> </a>
                </div>
                <div className="training-item">
                    <img
                        src="https://www.illuminancesolutions.com.au/wp-content/uploads/2017/09/Business-Foundations-logo.png"
                        alt=""
                    />
                    <h3>Business Foundation</h3>
                    <p>
                        Welcome to Business Foundation Course In this course we are providing
                        all required resource to You <small>Click Now</small>
                    </p>
                    <a href="./courses/Busl/busscon.html"><button>read more</button> </a>
                </div>
                <div className="training-item">
                    <img
                        src="https://ded9.com/wp-content/uploads/2021/04/digiwiseacademy-devops.jpeg"
                        alt=""
                        style={{ width: "50%" }}
                    />
                    <h3>DevOps for SDLC</h3>
                    <p>
                        Welcome to DevOps Course In this course we are providing all required
                        resource to You <small>Click Now</small>
                    </p>
                    <a href="./courses/devops/dev.html"><button>read more</button> </a>
                </div>
                <div className="training-item">
                    <img
                        src="https://ded9.com/wp-content/uploads/2021/04/digiwiseacademy-devops.jpeg"
                        alt=""
                        style={{ width: "50%" }}
                    />
                    <h3>DevOps for SDLC</h3>
                    <p>
                        Welcome to DevOps Course In this course we are providing all required
                        resource to You <small>Click Now</small>
                    </p>
                    <a href="./courses/devops/dev.html"><button>read more</button> </a>
                </div>
            </div>
        </div>

    )
}