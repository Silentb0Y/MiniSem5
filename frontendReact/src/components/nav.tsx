import "../all.css"
import "../home.css"
export function Nav () {
    return (
        <div>
            <header>
                <a href="#" className="logo">CodeForProgress</a>
                <nav>
                    <a href="#home" className="active">Home</a>
                    <a href="#About" >About</a>
                    <a href="#Services" >Services</a>
                    <a href="#Contact" >Contact</a>
                    <a href="#Portfolio" >Portfolio</a>
                </nav>
            </header>
        </div>
    )
}