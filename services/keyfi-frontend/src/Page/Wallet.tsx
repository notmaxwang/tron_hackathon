import { useState, useMemo } from 'react';
import type { WalletError } from '@tronweb3/tronwallet-abstract-adapter';
import { WalletDisconnectedError, WalletNotFoundError } from '@tronweb3/tronwallet-abstract-adapter';
import { useWallet, WalletProvider } from '@tronweb3/tronwallet-adapter-react-hooks';
import { WalletModalProvider } from '@tronweb3/tronwallet-adapter-react-ui';
import toast, {Toaster} from 'react-hot-toast';
import { tronWeb } from './tronweb';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField, Alert } from '@mui/material';
import { TronLinkAdapter } from '@tronweb3/tronwallet-adapter-tronlink';
import { WalletConnectAdapter } from '@tronweb3/tronwallet-adapter-walletconnect';
import {
    WalletActionButton,
    WalletConnectButton,
    WalletDisconnectButton,
    WalletSelectButton,
} from '@tronweb3/tronwallet-adapter-react-ui';
import { Button } from '@tronweb3/tronwallet-adapter-react-ui';
import './Wallet.css'


export default function Wallet() {
    function onError(e: WalletError) {
        if (e instanceof WalletNotFoundError || e instanceof WalletDisconnectedError) {
            toast.error(e.message);
        } else {
            toast.error(e.message);
        }
    }

    const adapters = useMemo(() => {
        const tronLinkAdapter = new TronLinkAdapter();
        const walletConnectAdapter = new WalletConnectAdapter({
            network: 'Nile',
            options: {
                relayUrl: 'wss://relay.walletconnect.com',
                projectId: '5fc507d8fc7ae913fff0b8071c7df231',
                metadata: {
                    name: 'Test DApp',
                    description: 'JustLend WalletConnect',
                    url: 'https://your-dapp-url.org/',
                    icons: ['https://your-dapp-url.org/mainLogo.svg'],
                },
            },
            web3ModalConfig: {
                themeMode: 'dark',
                themeVariables: {
                    '--w3m-z-index': '1000'
                },
            }
        });
        return [tronLinkAdapter, walletConnectAdapter];
    }, []);

    return (
        <WalletProvider onError={onError} autoConnect={true} disableAutoConnectOnLoad={true} adapters={adapters}>
            <WalletModalProvider>
                <UIComponent/>
                <Profile />
                <SignDemo />
            </WalletModalProvider>
            <Toaster />
        </WalletProvider>
    );
}

function UIComponent() {
    const rows = [
        { name: 'Connect Button', reactUI: WalletConnectButton },
        { name: 'Disconnect Button', reactUI: WalletDisconnectButton },
        { name: 'Action Button', reactUI: WalletActionButton },
        { name: 'Select Button', reactUI: WalletSelectButton },
    ];
    return (
        <div>
            <TableContainer style={{ overflow: 'visible' }} component="div">
                <Table sx={{  }} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell></TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rows.map((row) => (
                            <TableRow key={row.name} sx={{ '&:last-child td, &:last-child th': { border: 0 } }}>
                                <TableCell align='center'>
                                    <row.reactUI>{row.name}</row.reactUI>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    );
}

function Profile() {
    const { address, connected, wallet } = useWallet();
    return (
        <div>
            <h2>Wallet Connection Info</h2>
            <p>
                <span>Connection Status:</span> {connected ? 'Connected' : 'Disconnected'}
            </p>
            <p>
                <span>Your selected Wallet:</span> {wallet?.adapter.name}
            </p>
            <p>
                <span>Your Address:</span> {address}
            </p>
        </div>
    );
}

function SignDemo() {
  const { signMessage, signTransaction, address } = useWallet();
  const [message, setMessage] = useState('');
  const [signedMessage, setSignedMessage] = useState('');
  const receiver = 'TV3bpJQcY7DVQCNcEsKrDc1L4qEkSJvwrr';
  const [open, setOpen] = useState(false);

  async function onSignMessage() {
    // var str = "helloworld"; 
    // var HexStr = tronWeb.toHex(str);
    // const test = await tronWeb.trx.sign(HexStr);
    // console.log(test);
    // const res = await tronWeb.trx.sign(message);
    if(!message) {
        return(toast.error('Please type in a message to sign.'));
    }
      const res = await signMessage(message);
      setSignedMessage(res);
      console.log('signedmessage', res);
  }

  async function onSignTransaction() {
      const transaction = await tronWeb.transactionBuilder.sendTrx(receiver, tronWeb.toSun(0.001), address);

      const signedTransaction = await signTransaction(transaction);
    //   const signedTransaction = await tronWeb.trx.sign(transaction);
       await tronWeb.trx.sendRawTransaction(signedTransaction);
      setOpen(true);
  }
  return (
      <div style={{ marginBottom: 200 }}>
        <h2>Sign a message</h2>
        <p style={{ display: 'flex', justifyContent: 'center', flexWrap: 'wrap', wordBreak: 'break-all' }}>
            You can sign a message by clicking the button.
        </p>
        <TextField
            size="small"
            onChange={(e:any) => setMessage(e.target.value)}
            placeholder="Sign"
        ></TextField>
        <Button style={{ marginRight: '20px', }} onClick={onSignMessage}>
            Sign Message
        </Button>
        {signedMessage?<p>Your signed message is: {signedMessage}</p>:<></>}
        <h2>Sign a Transaction</h2>
        <p style={{ display: 'flex', justifyContent: 'center', flexWrap: 'wrap', wordBreak: 'break-all', margin: '15px' }}>
            You can transfer 0.001 Trx to &nbsp;<i>{receiver}</i>&nbsp;by click the button.
        </p>
        <Button onClick={onSignTransaction}>Transfer</Button>
        {open && (
            <Alert onClose={() => setOpen(false)} severity="success" sx={{ width: '100%', marginTop: 1 }}>
                Success! You can confirm your transfer on{' '}
                <a target="_blank" rel="noreferrer" href={`https://nile.tronscan.org/#/address/${address}`}>
                    Tron Scan
                </a>
            </Alert>
          )}
      </div>
  );
}
