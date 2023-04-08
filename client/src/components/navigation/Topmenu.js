import React from 'react';
import 'bulma/css/bulma.min.css';
import { Button, Navbar } from 'react-bulma-components';

class Topmenu extends React.Component {

    componentDidMount() {
    }

    render() {
        return (

            <Navbar {...this.props.args}>

                <Navbar.Brand>

                    <Navbar.Item renderAs="a" href="#">

                       <span className='logo'>Scriflow</span>

                    </Navbar.Item>

                    <Navbar.Burger />

                </Navbar.Brand>

                <Navbar.Menu>

                    <Navbar.Container>

                        <Navbar.Item hoverable={this.props.hoverable} active={this.props.itemActive} href="#">

                            <Navbar.Link arrowless={this.props.arrowless}>Files</Navbar.Link>

                            <Navbar.Dropdown up={this.props.up} right={this.props.right} boxed={this.props.boxed}>

                                <Navbar.Item href="#">Open Library</Navbar.Item>

                                <Navbar.Item>Local file</Navbar.Item>

                                <Navbar.Divider />

                                <Navbar.Item href="#">After divider</Navbar.Item>

                            </Navbar.Dropdown>

                        </Navbar.Item>

                        <Navbar.Item href="#">Second</Navbar.Item>

                    </Navbar.Container>

                    <Navbar.Container align="end">

                        <Navbar.Item href="#"><Button className="primary-bg">Log in</Button><Button color="secondary">Sign up</Button></Navbar.Item>
                        <Navbar.Item></Navbar.Item>

                    </Navbar.Container>

                </Navbar.Menu>

            </Navbar>

        );

    }
}

export default Topmenu;
