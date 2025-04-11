import React from 'react';
import { Form, Input, Button, Typography, Space, Divider } from 'antd';
import { UserOutlined, LockOutlined, EyeTwoTone, EyeInvisibleOutlined, GoogleOutlined, GithubOutlined } from '@ant-design/icons';
import './login.css';

const { Text, Link } = Typography;

// 自定义Twitter/X图标组件
const TwitterXIcon = () => (
  <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor">
    <path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z" />
  </svg>
);

const LoginPage: React.FC = () => {
  const [form] = Form.useForm();

  const onFinish = (values: any) => {
    console.log('Received values of form: ', values);
    // 这里添加登录逻辑
  };

  return (
    <div className="login-container">
      <div className="login-content">
        <div className="logo-container">
          <img 
            src="/src/assets/zoomeye-logo.svg" 
            alt="ZoomEye Logo" 
            className="zoomeye-logo" 
            onError={(e) => {
              // 如果logo图片加载失败，显示文字logo
              const target = e.target as HTMLImageElement;
              target.style.display = 'none';
              const parent = target.parentElement;
              if (parent) {
                const textLogo = document.createElement('div');
                textLogo.className = 'text-logo';
                textLogo.innerHTML = '<span style="color: #1890ff; font-weight: bold; font-size: 24px;">ZoomEye</span>';
                parent.appendChild(textLogo);
              }
            }}
          />
        </div>

        <div className="form-container">
          <h1 className="login-title">Login</h1>
          
          <Form
            form={form}
            name="login"
            className="login-form"
            initialValues={{ remember: true }}
            onFinish={onFinish}
            size="large"
          >
            <Form.Item
              name="username"
              rules={[{ required: true, message: 'Please enter your Email address / Username' }]}
            >
              <Input 
                prefix={<UserOutlined className="site-form-item-icon" />} 
                placeholder="Please enter Email address / Username" 
              />
            </Form.Item>

            <Form.Item
              name="password"
              rules={[{ required: true, message: 'Please enter password' }]}
            >
              <Input.Password
                prefix={<LockOutlined className="site-form-item-icon" />}
                placeholder="Please enter password"
                iconRender={visible => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)}
              />
            </Form.Item>

            <Form.Item>
              <Button type="primary" htmlType="submit" className="login-form-button">
                Login
              </Button>
            </Form.Item>

            <div className="terms-container">
              <Text className="terms-text">
                By Login, you agree to the 
                <Link href="#" className="terms-link">Zoomeye Service Agreement</Link>
                {' & '}
                <Link href="#" className="terms-link">User Privacy Policy</Link>
              </Text>
            </div>

            <div className="social-login">
              <Divider plain>Social Login:</Divider>
              <div className="social-icons">
                <Space size="middle">
                  <Button shape="circle" icon={<GoogleOutlined />} className="social-icon-button" />
                  <Button shape="circle" icon={<TwitterXIcon />} className="social-icon-button" />
                  <Button shape="circle" icon={<GithubOutlined />} className="social-icon-button" />
                </Space>
              </div>
            </div>

            <div className="bottom-links">
              <Space size="middle">
                <Link href="#" className="bottom-link">Create account</Link>
                <Link href="#" className="bottom-link">Forgot password?</Link>
              </Space>
            </div>
          </Form>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;