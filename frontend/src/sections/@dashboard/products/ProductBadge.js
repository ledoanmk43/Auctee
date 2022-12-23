import { useState, useEffect } from 'react';
import { redirect, useNavigate, useLocation, useSearchParams } from 'react-router-dom';

// material
import { useTheme } from '@mui/material/styles';
import useWebSocket, { ReadyState } from 'react-use-websocket';

import { Avatar, Stack, Badge, Typography, IconButton } from '@mui/material';
// component

// ----------------------------------------------------------------------

export default function ProductBadge({ auction, index, setLoaded }) {
  const [searchParams, setSearchParams] = useSearchParams();
  const id = searchParams.get('id');
  const [socketUrl, setSocketUrl] = useState('ws://localhost:1009/auctee/ws');

  const { lastMessage } = useWebSocket(socketUrl);

  const handleClick = (auction) => {
    window.location.assign(`/auctee/auction/detail?id=${auction.Id}&product=${auction.product_id}`);
    window.localStorage.setItem(`NOTI_${auction.Id}`, false);
  };
  const notificationsLabel = (count) => {
    if (count === 0) {
      return 'no notifications';
    }
    if (count > 99) {
      return 'more than 99 notifications';
    }
    return `${count} notifications`;
  };
  const [notification, setNotification] = useState();
  const [body, setBody] = useState();

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    const body = lastMessage && JSON.parse(lastMessage.data);
    // eslint-disable-next-line no-unused-expressions
    if (id !== String(body?.auction_id) && auction.Id === body?.auction_id) {
      window.localStorage.setItem(`NOTI_${body?.auction_id}`, true);
      setBody(lastMessage && JSON.parse(lastMessage.data));
    }
  }, [lastMessage]);

  useEffect(() => {
    // eslint-disable-next-line no-unused-expressions
    if (id !== String(body?.auction_id) || id !== String(auction.Id)) {
      setNotification(window.localStorage.getItem(`NOTI_${body?.auction_id}`));
    }
  }, [body]);
  useEffect(() => {
    setNotification(window.localStorage.getItem(`NOTI_${auction.Id}`));
    setLoaded(true);
  }, []);

  return (
    <Stack
      key={index}
      onClick={() => handleClick(auction)}
      sx={{
        mt: 1,
        position: 'relative',
      }}
    >
      <IconButton aria-label={notificationsLabel(100)}>
        <Badge overlap="circular" color="error" badgeContent={notification === 'true' ? '!' : 0}>
          <Avatar
            alt="Remy Sharp"
            src={auction.image_path}
            sx={{ width: 55, height: 55, '&:hover': { opacity: 0.72 } }}
          />
        </Badge>
      </IconButton>
    </Stack>
  );
}
