import { createPlugin } from '@backstage/core';

import WelcomePage from './components/WelcomePage';

import Adddata from './components/Adddata';
import afterlogin from './components/afterlogin';


export const plugin = createPlugin({

  id: 'welcome',

  register({ router }) {

    router.registerRoute('/', WelcomePage);
    router.registerRoute('/adddata', Adddata);
    router.registerRoute('/afterlogin', afterlogin);

  },

});