#include <signal.h>
#include "IRCClient.h"

volatile bool running;

void signalHandler(int signal)
{
    std::cout << "Stop!" << std::endl;
    running = false;
}

void onAssertJoin(IRCMessage message, IRCClient* client)
{
    if (message.prefix.nick != "assert") {
        return;
    }

    client->SendIRC("PRIVMSG #grunndev :Hello assert!");
}

int main(int argc, char const *argv[])
{
    IRCClient client;
    std::string nick("hello_assert_bot");
    std::string user("hello_assert_bot");

    client.HookIRCCommand("JOIN", &onAssertJoin);

    if (!client.InitSocket()) {
        return 1;
    }

    if (!client.Connect("chat.freenode.net", 6665)) {
        return 1;
    }

    if (client.Login(nick, user))
    {
        running = true;
        signal(SIGINT, signalHandler);

        client.SendIRC("JOIN #grunndev");

        while (client.Connected() && running) {
            client.ReceiveData();
        }
    }

    if (client.Connected()) {
        client.Disconnect();
    }

    return 0;
}
