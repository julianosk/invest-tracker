import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import AssetAllocation from './pages/AssetAllocation';
import InvestmentList from './components/InvestmentList';

const App: React.FC = () => {
  return (
    <Router>
      <div>
        <h1>Investment Management App</h1>
        <Switch>
          <Route path="/" exact component={InvestmentList} />
          <Route path="/asset-allocation" component={AssetAllocation} />
        </Switch>
      </div>
    </Router>
  );
};

export default App;