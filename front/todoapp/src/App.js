import { BrowserRouter, Route, Routes } from "react-router-dom";
import Signin from "./components/auth/Signin";
import Signup from "./components/auth/Signup";
import Top from "./components/views/Top";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Top />} />
                <Route path="/signup" element={<Signup />} />
                <Route path="/signin" element={<Signin />} />
            </Routes>
        </BrowserRouter>
    )
}

export default App;
