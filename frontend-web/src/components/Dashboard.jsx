import {useNavigate} from "react-router-dom";

export const Dashboard = () => {
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate("/login");
    };

    return (
        <div className={"card"}>
            <h1>secret page</h1>
            <p>successful auth</p>
            <button onClick={handleLogout} style={{marginTop: '20px', backgroundColor: '#e74c3c'}}>
                Log out
            </button>
        </div>
    );
}