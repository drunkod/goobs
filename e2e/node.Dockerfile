FROM node:12.22-alpine

# env ssh_pub_key=""



RUN mkdir -p /root/.ssh \
    && chmod 0700 /root/.ssh \
    && passwd -u root \
    && echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDYSC/8fD55n+0LuaDAEF33vkEUfGFkY+qj/lWslFf22zlKzELhbepRy95QYfCI5kjA4cH3qEcWZCRJ6hShBChooU8WInrPlTiGN3hZg7hGtHST7MXHIgjVqFXvFn0+CSTCxYkJ4x7UDXhNmMl/4XoNXaCrsRrav3lUfSm7c+WDHTaqC3kdOYLnrboDyxnTbSKM7C+FAeShas1ZE6PD20InzodUDJTYnzjs+p5OepRUKG5TNHolF0emNWgux8mhDFLPb5mmOo6aquVIaXmf2TwOH3YHuRnoCRduncIwK+N2r4AtmJP3kv6QAOMSw0BRIankx/JJniYFOYheU7Q++SNQKcvUwf7h3LKAD8KcqEQtjD5J6oHQXex3T8lzNiQSw2HSsrW/XoOKME7hwCGIOmno6O8pRWoJkY0OcFPrq4lwqv8lzrJq+znC48XsTJr562ctx5ML4Pv6n1fYVvY7RSZX29WPwY/JJuCsg2lpOf1opOiZ1uBVldXR2it4nwWI/P8= alex@fedora" > /root/.ssh/authorized_keys \
    && apk add openrc openssh \
    && ssh-keygen -A \
    && echo -e "PasswordAuthentication no" >> /etc/ssh/sshd_config \
    && mkdir -p /run/openrc \
    && touch /run/openrc/softlevel

# WORKDIR /app

# COPY . .

# RUN yarn

# RUN yarn build

ENTRYPOINT ["sh", "-c", "rc-status; rc-service sshd start"]