import { Navigate } from "react-router-dom";
import { useAuthContext } from "../../context/AuthContext";
import Header from "./Header";
import TodoList from "./TodoList";

const Top = () => {
    const info = useAuthContext();

    if (!info.user) {
        return <Navigate to="/signin"/>
    } else {
        return (
            <>
                <Header />
                <TodoList />
            </>
        );
    }
}

export default Top;
