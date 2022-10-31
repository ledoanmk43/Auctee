import { Navigate, useRoutes } from 'react-router-dom';
// layouts
import MainLayout from './layouts/main';
import LogoOnlyLayout from './layouts/LogoOnlyLayout';
//
import Blog from './pages/Blog';
import User from './pages/User';
import Address from './pages/Address';
import ChangePwd from './pages/ChangePwd';
import Inventory from './pages/Inventory';

import Purchase from './pages/Purchase';
import Login from './pages/Login';
import NotFound from './pages/Page404';
import Register from './pages/Register';
import Home from './pages/Home';
import ProductDetail from './pages/ProductDetail';
import SearchProduct from './pages/SearchProduct';

import DashboardApp from './pages/DashboardApp';

// ----------------------------------------------------------------------

export default function Router() {
  return useRoutes([
    {
      path: '/auctee',
      element: <MainLayout />,
      children: [
        { path: 'home', element: <Home /> },
        { path: 'auction/detail', element: <ProductDetail /> },
        { path: 'search', element: <SearchProduct /> },
        { path: 'user/profile', element: <User /> },
        { path: 'user/change-password', element: <ChangePwd /> },
        { path: 'user/address', element: <Address /> },
        { path: 'user/inventory', element: <Inventory /> },
        { path: 'user/purchase', element: <Purchase /> },
        { path: 'blog', element: <Blog /> },
      ],
    },
    {
      path: '/auctee',
      children: [
        {
          path: 'login',
          element: <Login />,
        },
        {
          path: 'register',
          element: <Register />,
        },
      ],
    },
    {
      path: '/',
      element: <LogoOnlyLayout />,
      children: [
        { path: '/', element: <Navigate to="/auctee/home" /> },
        { path: '404', element: <NotFound /> },
        { path: '*', element: <Navigate to="/404" /> },
      ],
    },
    {
      path: '*',
      element: <Navigate to="/404" replace />,
    },
  ]);
}
