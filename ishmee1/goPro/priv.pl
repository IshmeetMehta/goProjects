#!/usr/bin/perl

use strict;
use warnings;
use JSON::PP;
use Crypt::PK::RSA;

my $filename = 'service-account-key.pem';
my $pk = Crypt::PK::RSA->new($filename);
my $private_key = $pk->export_key_jwk('private');

print "$private_key";
