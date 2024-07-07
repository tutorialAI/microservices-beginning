FROM php:8.1.0-fpm
RUN apt-get update \
    && php -r "readfile('https://getcomposer.org/installer');" | php -- --install-dir=/usr/bin/ --filename=composer
USER 0
RUN apt-get install -y librdkafka-dev \
    && pecl channel-update pecl.php.net && pecl install rdkafka
USER $CONTAINER_USER_ID