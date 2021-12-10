import React from 'react';
import 'bulma/css/bulma.min.css';
import {
    Modal,
    ModalCard,
    ModalCardBody,
    ModalCardHead,
    ModalCardFoot, ModalCardTitle,
    Content, Button, MediaItem,
    Media, Image
} from 'react-bulma-components';

class LoadFileModal extends React.Component {

    constructor(prop) {
        super(prop)
        this.state = { show: false };
    }
    
    componentDidMount() {

    }

    

    render() {
        const showHideClassName = this.state.show ? "modal is-active" : "modal";
          
        return (<Modal
            onClose={function noRefCheck() { }}
            showClose={false}
            className={showHideClassName}
        >
            <ModalCard>
                <ModalCardHead>
                    <ModalCardTitle>
                        Title
                    </ModalCardTitle>
                </ModalCardHead>
                <ModalCardBody>
                    <Media>
                        <MediaItem
                            align="left"
                            renderAs="figure"
                        >
                            <Image
                                alt="64x64"
                                size={64}
                                src="http://bulma.io/images/placeholders/128x128.png"
                            />
                        </MediaItem>
                        <MediaItem>
                            <Content>
                                <p>
                                    <strong>
                                        John Smith
                                    </strong>
                                    {' '}
                                    <small>
                                        @johnsmith
                                    </small>
                                    {' '}
                                    <small>
                                        31m
                                    </small>
                                    <br />
                                    If the children of the Modal is a card, the close button will be on the Card Head instead than the top-right corner You can also pass showClose = false to Card.Head to hide the close button
                                </p>
                                <Button>Test</Button>
                            </Content>
                        </MediaItem>
                    </Media>
                </ModalCardBody>
                <ModalCardFoot
                    align="right"
                    hasAddons
                    renderAs={function noRefCheck() { }}
                >
                </ModalCardFoot>
            </ModalCard>
        </Modal>)

    }
}

export default LoadFileModal;