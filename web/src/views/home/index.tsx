import React from 'react';
import { Typography, Card, Row, Col, Statistic } from 'antd';
import { UserOutlined, DatabaseOutlined, ApiOutlined, SettingOutlined } from '@ant-design/icons';

const { Title, Paragraph } = Typography;

const HomePage: React.FC = () => {
  return (
    <div style={{ padding: '24px' }}>
      <Typography>
        <Title level={2}>Dracula 系统控制台</Title>
        <Paragraph>
          欢迎使用Dracula系统，这是一个强大的网络安全管理平台。
        </Paragraph>
      </Typography>

      <Row gutter={16} style={{ marginTop: '24px' }}>
        <Col span={6}>
          <Card>
            <Statistic
              title="用户数量"
              value={42}
              prefix={<UserOutlined />}
            />
          </Card>
        </Col>
        <Col span={6}>
          <Card>
            <Statistic
              title="数据库连接"
              value={12}
              prefix={<DatabaseOutlined />}
            />
          </Card>
        </Col>
        <Col span={6}>
          <Card>
            <Statistic
              title="API调用"
              value={1024}
              prefix={<ApiOutlined />}
            />
          </Card>
        </Col>
        <Col span={6}>
          <Card>
            <Statistic
              title="系统设置"
              value={8}
              prefix={<SettingOutlined />}
            />
          </Card>
        </Col>
      </Row>

      <Card style={{ marginTop: '24px' }}>
        <Title level={4}>系统概览</Title>
        <Paragraph>
          Dracula系统提供了全面的网络安全管理功能，包括用户管理、数据分析、安全监控等。
          通过直观的界面，您可以轻松管理和监控系统的各个方面。
        </Paragraph>
      </Card>
    </div>
  );
};

export default HomePage;