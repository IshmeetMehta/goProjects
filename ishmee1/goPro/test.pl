#!/usr/bin/perl

use strict;
use warnings;
use JSON::PP;
use Crypt::PK::RSA;

my $filename = 'service-account.pem';
my $pk = Crypt::PK::RSA->new($filename);
print $pk;
my $public_key = $pk->export_key_jwk('public');

print "$public_key";
