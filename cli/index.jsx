import dva from 'dva';
import { browserHistory } from 'dva/router'
import './static/index.css';

// 1. Initialize
const app = dva({
    history: browserHistory
});

// 2. Plugins
// app.use({});

// 3. Model
app.model(require('./models/index'));

// 4. Router
// import RouterConfig from './router'
app.router(require('./router'));

// 5. Start
app.start('#root');
