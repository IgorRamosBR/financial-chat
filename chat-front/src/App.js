import './App.css';
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Login from "./login/Login";
import Register from "./register/Register";
import Chat from "./chat/Chat";

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route exact path="/" component={Login} />
          <Route exact path="/register" component={Register} />
          <Route exact path="/chat" component={Chat} />
        </Switch>
      </Router>
    </div>
  );
}

export default App;
