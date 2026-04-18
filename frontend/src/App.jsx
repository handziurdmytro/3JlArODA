import {Auth} from './components/Auth/Auth.jsx'
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import {ProtectedRoute} from "./components/Auth/ProtectedRoute.jsx";
import { MainPage } from './components/MainPage/MainPage.jsx';

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
                                <MainPage/>
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
