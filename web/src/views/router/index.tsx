import { createBrowserRouter, RouterProvider, Navigate, Outlet } from 'react-router-dom';
import { useEffect, useState } from 'react';
import LoginPage from '../login/login';

// 由于Home组件尚未创建，先创建一个简单的Home组件
const Home = () => {
  return (
    <div style={{ padding: '20px' }}>
      <h1>Dracula 系统主页</h1>
      <p>欢迎使用Dracula系统</p>
    </div>
  );
};

// 路由守卫组件，用于保护需要登录才能访问的路由
const ProtectedRoute = () => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    // 检查用户是否已登录
    // 这里可以从localStorage或cookie中获取token，或者调用API验证token
    const token = localStorage.getItem('token');
    setIsAuthenticated(!!token);
    setIsLoading(false);
  }, []);

  if (isLoading) {
    return <div>加载中...</div>;
  }

  return isAuthenticated ? <Outlet /> : <Navigate to="/login" replace />;
};

// 创建路由配置
const router = createBrowserRouter([
  {
    path: '/',
    element: <Navigate to="/home" replace />,
  },
  {
    path: '/login',
    element: <LoginPage />,
  },
  {
    path: '/',
    element: <ProtectedRoute />,
    children: [
      {
        path: 'home',
        element: <Home />,
      },
      // 可以在这里添加更多需要登录才能访问的路由
    ],
  },
  {
    path: '*',
    element: <div>404 - 页面不存在</div>,
  },
]);

// 路由提供者组件
const AppRouter = () => {
  return <RouterProvider router={router} />;
};

export default AppRouter;