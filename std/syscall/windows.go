// windows系统调用
//
// WSARecv
// https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-wsarecv
//
// WSAECONNRESET 10054
// For a stream socket, the virtual circuit was reset by the remote side.
// The application should close the socket as it is no longer usable.
// An existing connection was forcibly closed by the remote host.
// https://learn.microsoft.com/en-us/windows/win32/winsock/windows-sockets-error-codes-2
package syscall
