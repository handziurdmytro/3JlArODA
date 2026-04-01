import {Auth} from './components/Auth.jsx'
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import {ProtectedRoute} from "./components/ProtectedRoute.jsx";
import {Dashboard} from "./components/Dashboard.jsx";

function App() {
    return (
        <BrowserRouter>
            <div id={"root"}>
                <Routes>
                    <Route path={'/login'} element={<Auth/>}/>
                    <Route
                        path={'/dashboard'}
                        element={
                            <ProtectedRoute>
                                <Dashboard/>
                            </ProtectedRoute>
                        }
                    />
                    <Route path="*" element={<Navigate to="/dashboard" replace/>}/>
                </Routes>
            </div>
        </BrowserRouter>
    );
}

export default App
