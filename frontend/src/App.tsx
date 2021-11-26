import { BrowserRouter, Route } from "react-router-dom";
import { Login } from "@modules/Auth/Login";
import { Register } from "@modules/Auth/Register";
import { ProtectedRoute } from "@components/ProtectedRoute/ProtectedRoute";
import { Homepage } from "@modules/Homepage/Homepage";

const App = () => {
  return (
    <div className="App">
      <BrowserRouter>
        <Route path={"/auth/login"}>
          <Login />
        </Route>
        <Route path={"/auth/register"}>
          <Register />
        </Route>

        <ProtectedRoute
          path={"/"}
        ><Homepage/></ProtectedRoute>
      </BrowserRouter>
    </div>
  );
};

export default App;
